# GoShell

A lightweight shell implemented in Go that supports basic file system operations and the execution of external commands. The shell mimics the functionality of popular terminal applications and includes commands like `ls`, `cd`, `pwd`, `touch`, `mkdir`, `rm`, `cat`, `echo`, `clear`, `cp`, `mv`, `chmod`, `stat`, and more.

## Features

- **File System Operations**:

  - List files (`ls`)
  - Change directories (`cd`)
  - Display current directory (`pwd`)
  - Create files (`touch`) and directories (`mkdir`)
  - Remove files or directories (`rm`)
  - Copy files or directories (`cp`)
  - Move or rename files/directories (`mv`)
  - View file content (`cat`)

- **Utility Commands**:

  - Display a message (`echo`)
  - Clear the terminal screen (`clear`)
  - Display file or directory stats (`stat`)
  - Change file permissions (`chmod`)
  - Manual pages for commands (`man`)

- **External Commands**:

  - Supports running any external command not explicitly defined (e.g., `go`, `python`, `git`).

- **Interactive Shell Loop**:
  - Continuous command execution with a user-friendly prompt showing the current directory.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/<your-username>/goshell.git
   cd goshell
   ```

2. Build the project:

   ```bash
   go build -o goshell
   ```

3. Run the shell:
   ```bash
   ./goshell
   ```

## Usage

After running the shell, you can execute various commands:

### Examples:

- List files in the current directory:

  ```bash
  ls
  ```

- Change to a specific directory:

  ```bash
  cd folder_name
  ```

- Create a file:

  ```bash
  touch new_file.txt
  ```

- Create a directory:

  ```bash
  mkdir new_folder
  ```

- Remove a file:

  ```bash
  rm file_to_delete.txt
  ```

- Copy a file:

  ```bash
  cp source.txt destination.txt
  ```

- View file content:

  ```bash
  cat file.txt
  ```

- Run an external command (e.g., `go version`):
  ```bash
  go version
  ```

## Command Reference

| Command | Description                                           |
| ------- | ----------------------------------------------------- |
| `ls`    | Lists files and directories in the current directory. |
| `cd`    | Changes the current directory.                        |
| `pwd`   | Prints the current working directory.                 |
| `touch` | Creates a new empty file.                             |
| `mkdir` | Creates a new directory.                              |
| `rm`    | Removes a file or directory. Use `/f` for force.      |
| `cat`   | Displays the content of a file.                       |
| `echo`  | Prints a message to the screen.                       |
| `clear` | Clears the terminal screen.                           |
| `cp`    | Copies a file or directory.                           |
| `mv`    | Moves or renames a file or directory.                 |
| `chmod` | Changes the permissions of a file.                    |
| `stat`  | Displays detailed information about a file or dir.    |
| `man`   | Displays a brief manual for a command.                |
| `exit`  | Exits the shell.                                      |

## Project Structure

- `main.go`: Contains the implementation of the shell.
- `Commands`: Implemented as switch cases for easy extensibility.

## Future Enhancements

- Add support for pipes (`|`) and redirections (`>`, `>>`, `<`).
- Implement autocomplete for commands and paths.
- Support background tasks with `&`.

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
