package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func WalkFunction(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fileInfo.IsDir() == false {
		return nil
	}
	fmt.Println(path)
	return nil
}

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("not enough arguments!")
		os.Exit(1)
	}
	Path := arguments[1]
	err := filepath.Walk(Path, WalkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
