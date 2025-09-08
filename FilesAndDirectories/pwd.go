package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args

	// Getwd returns an absolute path name corresponding to the current directory.
	//  If the current directory can be reached via multiple paths (due to symbolic links),
	// On Unix platforms, if the environment variable PWD provides an absolute name,
	// and it is a name of the current directory, it is returned.
	// Getwd may return any one of them.
	pwd, err := os.Getwd()

	if err == nil {
		fmt.Println(pwd)
	} else {
		fmt.Println("err:", err)
	}

	if len(arguments) == 1 || arguments[1] != "-P" {
		return
	}

	// checking is this is a symbolic link or not
	fileInfo, err := os.Lstat(pwd)

	if fileInfo.Mode()&os.ModeSymlink != 0 {
		realPath, err := filepath.EvalSymlinks(pwd)
		if err == nil {
			fmt.Println(realPath)
		}
	}

}
