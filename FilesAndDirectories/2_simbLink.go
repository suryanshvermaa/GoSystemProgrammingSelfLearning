package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument")
		os.Exit(1)
	}
	filename := arguments[1]

	// Lstat returns a [FileInfo] describing the named file.
	// If the file is a symbolic link, the returned FileInfo describes
	//  the symbolic link. Lstat makes no attempt to follow the link.
	// If there is an error, it will be of type [*PathError].
	fileInfo, err := os.Lstat(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if fileInfo.Mode()&os.ModeSymlink != 0 {
		fmt.Println(filename, "is a symbolic link")
		realPath, err := filepath.EvalSymlinks(filename)
		if err == nil {
			fmt.Println("path:", realPath)
		} else {
			fmt.Println("Error resolving symbolic link:", err)
		}
	} else {
		fmt.Println(filename, "is not a symbolic link")
	}
}
