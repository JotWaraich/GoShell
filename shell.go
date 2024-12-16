package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// Run the shell loop
	runShell()
}

// Run the shell with a loop for commands
func runShell() {
	// Start at the current working directory
	currentDir, _ := os.Getwd()

	// Main loop
	for {
		// Display the prompt with the current directory
		fmt.Printf("%s$ ", currentDir)

		// Read user input
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Handle commands
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "ls":
			listFiles(currentDir)
		case "cd":
			if len(args) > 1 {
				currentDir = changeDirectory(args[1], currentDir)
			} else {
				fmt.Println("Error: 'cd' command requires a path argument.")
			}
		case "pwd":
			printWorkingDirectory(currentDir)
		case "touch":
			if len(args) > 1 {
				createFile(args[1], currentDir)
			} else {
				fmt.Println("Error: 'touch' command requires a file name.")
			}
		case "mkdir":
			if len(args) > 1 {
				createDirectory(args[1], currentDir)
			} else {
				fmt.Println("Error: 'mkdir' command requires a directory name.")
			}
		case "rm":
			if len(args) > 1 {
				forceDelete := false
				if len(args) > 2 && args[1] == "/f" {
					forceDelete = true
					removeFileOrDirectory(args[2], currentDir, forceDelete)
				} else {
					removeFileOrDirectory(args[1], currentDir, forceDelete)
				}
			} else {
				fmt.Println("Error: 'rm' command requires a file or directory name.")
			}
		case "cat":
			if len(args) > 1 {
				displayFileContent(args[1], currentDir)
			} else {
				fmt.Println("Error: 'cat' command requires a file name.")
			}
		case "echo":
			if len(args) > 1 {
				echoMessage(strings.Join(args[1:], " "))
			} else {
				fmt.Println("Error: 'echo' command requires a message.")
			}
		case "clear":
			clearScreen()
		case "cp":
			if len(args) > 2 {
				copyFileOrDirectory(args[1], args[2], currentDir)
			} else {
				fmt.Println("Error: 'cp' command requires source and destination.")
			}
		case "mv":
			if len(args) > 2 {
				moveOrRenameFileOrDirectory(args[1], args[2], currentDir)
			} else {
				fmt.Println("Error: 'mv' command requires source and destination.")
			}
		case "man":
			if len(args) > 1 {
				showManPage(args[1])
			} else {
				fmt.Println("Error: 'man' command requires a command to display help for.")
			}
		case "chmod":
			if len(args) > 2 {
				permissions, err := strconv.ParseInt(args[1], 8, 32)
				if err != nil {
					fmt.Println("Error: Invalid permission format.")
					break
				}
				changeFilePermissions(args[2], currentDir, os.FileMode(permissions))
			} else {
				fmt.Println("Error: 'chmod' command requires permissions and file name.")
			}
		case "stat":
			if len(args) > 1 {
				showFileInfo(args[1], currentDir)
			} else {
				fmt.Println("Error: 'stat' command requires a file or directory name.")
			}
		case "exit":
			fmt.Println("Exiting Shell...")
			return
		default:
			runExternalCommand(args[0], args[1:]...) // Run any other command
		}
	}
}

// List files in the current directory
func listFiles(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("[DIR]  ", file.Name())
		} else {
			fmt.Println(file.Name())
		}
	}
}

// Change to a specified directory
func changeDirectory(newDir, currentDir string) string {
	// Resolve the absolute path
	absolutePath, err := filepath.Abs(filepath.Join(currentDir, newDir))
	if err != nil {
		fmt.Println("Error resolving path:", err)
		return currentDir
	}

	// Check if the directory exists
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		fmt.Println("Error: Directory does not exist:", newDir)
		return currentDir
	}

	// Change to the new directory
	err = os.Chdir(absolutePath)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return currentDir
	}

	// Return the new current directory
	return absolutePath
}

// Print the current directory
func printWorkingDirectory(currentDir string) {
	fmt.Println(currentDir)
}

// Create a new file
func createFile(fileName, currentDir string) {
	filePath := filepath.Join(currentDir, fileName)
	// Create the file and ensure it is closed after creation
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after creation

	fmt.Println("File created:", filePath)
}

// Create a new directory
func createDirectory(dirName, currentDir string) {
	dirPath := filepath.Join(currentDir, dirName)
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	fmt.Println("Directory created:", dirPath)
}

// Remove a file or directory (with optional force delete)
func removeFileOrDirectory(name, currentDir string, forceDelete bool) {
	path := filepath.Join(currentDir, name)

	// Check if the path exists
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println("Error: File or directory does not exist:", name)
		return
	}

	// Check if force delete is enabled
	if forceDelete {
		err = os.RemoveAll(path) // Force delete for both files and directories
		if err != nil {
			fmt.Println("Error force deleting file or directory:", err)
			return
		}
		fmt.Println("Force deleted:", path)
		return
	}

	// Remove directory or file normally
	if info.IsDir() {
		err = os.Remove(path) // Remove only empty directories
		if err != nil {
			fmt.Println("Error removing directory:", err)
			return
		}
		fmt.Println("Directory removed:", path)
	} else {
		err = os.Remove(path) // Remove single file
		if err != nil {
			fmt.Println("Error removing file:", err)
			return
		}
		fmt.Println("File removed:", path)
	}
}

func displayFileContent(filename, currentDir string) {
	filePath := filepath.Join(currentDir, filename)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(content))
}

func echoMessage(message string) {
	fmt.Println(message)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// Copy a file or directory
func copyFileOrDirectory(src, dest, currentDir string) {
    srcPath := filepath.Join(currentDir, src)
    destPath := filepath.Join(currentDir, dest)

    srcInfo, err := os.Stat(srcPath)
    if err != nil {
        fmt.Println("Error: Source file or directory does not exist:", src)
        return
    }

    if srcInfo.IsDir() {
        err := copyDirectory(srcPath, destPath)
        if err != nil {
            fmt.Println("Error copying directory:", err)
        } else {
            fmt.Println("Directory copied:", src, "to", dest)
        }
    } else {
        err := copyFile(srcPath, destPath)
        if err != nil {
            fmt.Println("Error copying file:", err)
        } else {
            fmt.Println("File copied:", src, "to", dest)
        }
    }
}

// Copy a file
func copyFile(src, dest string) error {
    input, err := ioutil.ReadFile(src)
    if err != nil {
        return err
    }
    err = ioutil.WriteFile(dest, input, 0644)
    return err
}

// Copy a directory recursively
func copyDirectory(src, dest string) error {
    err := os.MkdirAll(dest, 0755)
    if err != nil {
        return err
    }
    files, err := ioutil.ReadDir(src)
    if err != nil {
        return err
    }
    for _, file := range files {
        srcPath := filepath.Join(src, file.Name())
        destPath := filepath.Join(dest, file.Name())
        if file.IsDir() {
            err = copyDirectory(srcPath, destPath)
        } else {
            err = copyFile(srcPath, destPath)
        }
        if err != nil {
            return err
        }
    }
    return nil
}

// Move or rename a file or directory
func moveOrRenameFileOrDirectory(src, dest, currentDir string) {
    srcPath := filepath.Join(currentDir, src)
    destPath := filepath.Join(currentDir, dest)

    err := os.Rename(srcPath, destPath)
    if err != nil {
        fmt.Println("Error moving or renaming:", err)
    } else {
        fmt.Println("Moved/Renamed:", src, "to", dest)
    }
}

func showManPage(command string) {
    switch command {
    case "ls":
        fmt.Println("ls: List directory contents")
    case "cd":
        fmt.Println("cd: Change the current directory")
    case "touch":
        fmt.Println("touch: Create a new empty file")
    case "mkdir":
        fmt.Println("mkdir: Create a new directory")
    case "rm":
        fmt.Println("rm: Remove files or directories")
    case "cp":
        fmt.Println("cp: Copy files or directories")
    case "mv":
        fmt.Println("mv: Move or rename files or directories")
    case "echo":
        fmt.Println("echo: Display a message")
    case "clear":
        fmt.Println("clear: Clear the terminal screen")
    default:
        fmt.Println("No manual entry for:", command)
    }
}

// Change file permissions
func changeFilePermissions(filename, currentDir string, permissions os.FileMode) {
    filePath := filepath.Join(currentDir, filename)
    err := os.Chmod(filePath, permissions)
    if err != nil {
        fmt.Println("Error changing permissions:", err)
        return
    }
    fmt.Println("Permissions changed for:", filePath)
}

// Display file/directory status (info)
func showFileInfo(filename, currentDir string) {
    filePath := filepath.Join(currentDir, filename)
    info, err := os.Stat(filePath)
    if err != nil {
        fmt.Println("Error retrieving file info:", err)
        return
    }
    fmt.Println("Name:", info.Name())
    fmt.Println("Size:", info.Size(), "bytes")
    fmt.Println("Permissions:", info.Mode())
    fmt.Println("Is Directory:", info.IsDir())
    fmt.Println("Last Modified:", info.ModTime())
}

// Run an external command (git, code, etc.)
func runExternalCommand(command string, args ...string) {
    cmd := exec.Command(command, args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Printf("Error executing command '%s': %v\n", command, err)
    }
}