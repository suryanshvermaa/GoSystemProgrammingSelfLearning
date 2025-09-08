package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func WalkFunction(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fileInfo.Mode().IsDir() || fileInfo.Mode().IsRegular() {
		fmt.Println(path)
	}
	return nil
}

func main() {
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("not enough arguments!")z
		os.Exit(1)
	}
	Path := flags[0]
	err := filepath.Walk(Path, WalkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
