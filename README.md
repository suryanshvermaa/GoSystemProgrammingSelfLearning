# GoSystemProgrammingSelfLearning

## Description

> This repository is for learning system programming using Golang. It contains various examples and utilities to explore file operations, directory traversal, symbolic links, and more.

---

### Prerequisites

- `Go` >= 1.18

---

## Installation

Clone the repository:

```bash
git clone https://github.com/suryanshvermaa/GoSystemProgrammingSelfLearning.git

cd GoSystemProgrammingSelfLearning
```

---

## Project Structure

```
GoSystemProgrammingSelfLearning/
├── README.md
├── FilesAndDirectories/
│   ├── 1_flagPackege.go                # Example of using the flag package
│   ├── 2_simbLink.go                   # Check if a file is a symbolic link
│   ├── 3_pwd.go                        # Implementation of the `pwd` command
│   ├── 4_which.go                      # Implementation of the `which` command
│   ├── 5_PermissionBitsOfFile.go       # Display permission bits of a file
│   ├── 6_traversing_diretoryTree.go    # Traverse a directory tree
│   ├── 7_visitingOnlyDirectories.go    # Visit only directories in a tree
│   ├── 8_findv1.go                     # Implementation of a basic `find` command
│   └── FileOperation/
│       ├── deleteFile.go               # Delete a file
│       ├── dummy.json                  # Sample JSON file
│       ├── renameFile.go               # Rename a file
```

---

## File Descriptions

### Top-Level Files
- **README.md**: This file, providing an overview of the project.

### FilesAndDirectories
- **1_flagPackege.go**: Demonstrates the use of the `flag` package for parsing command-line arguments.
- **2_simbLink.go**: Checks if a file is a symbolic link and resolves its real path.
- **3_pwd.go**: Implements the `pwd` command to print the current working directory.
- **4_which.go**: Implements the `which` command to locate a command in the system's PATH.
- **5_PermissionBitsOfFile.go**: Displays the permission bits of a specified file.
- **6_traversing_diretoryTree.go**: Traverses a directory tree and prints all paths.
- **7_visitingOnlyDirectories.go**: Visits and prints only directories in a directory tree.
- **8_findv1.go**: A basic implementation of the `find` command to locate files and directories.

### FilesAndDirectories/FileOperation
- **deleteFile.go**: Deletes a specified file.
- **dummy.json**: A sample JSON file for testing purposes.
- **renameFile.go**: Renames a file, with support for overwriting.

---

## Usage

Each Go file in the project is a standalone program. You can run them using the `go run` command. For example:

```bash
go run FilesAndDirectories/3_pwd.go
```

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.