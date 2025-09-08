package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	minusOverwrite := flag.Bool("overwrite", false, "overwrite")
	flag.Parse()

	flags := flag.Args()

	if len(flags) < 2 {
		fmt.Println("please privide two arguments")
		os.Exit(1)
	}

	source := flags[0]
	dest := flags[1]

	fileInfo, err := os.Stat(source)
	if err == nil {
		mode := fileInfo.Mode()
		if mode.IsRegular() == false {
			fmt.Println("sorry we only support regular files")
			os.Exit(1)
		}
	} else {
		fmt.Println("error in reading source")
	}
	newDest := dest
	destInfo, err := os.Stat(dest)
	if err == nil {
		mode := destInfo.Mode()
		if mode.IsDir() {
			justTheName := filepath.Base(source)
			newDest = dest + "/" + justTheName
		}
	}

	dest = newDest
	destInfo, err = os.Stat(dest)
	if err == nil {
		if *minusOverwrite == false {
			fmt.Println("destination file already exists!")
			os.Exit(1)
		}
	}

	err = os.Rename(source, dest)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
