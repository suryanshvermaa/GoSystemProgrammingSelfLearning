package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// which command returns first path mathed path from $PATH
// Implementing which command with -s and -a options.
// If -a is present then return all paths otherwise only first one.
func main() {
	minusA := flag.Bool("a", false, "a")
	minusS := flag.Bool("s", false, "s")

	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("please provide an argument")
		os.Exit(1)
	}

	file := flags[0]

	FoundIt := false
	path := os.Getenv("PATH")
	pathSlice := strings.Split(path, ":")

	for _, dir := range pathSlice {
		fullPath := dir + "/" + file

		// Stat returns a [FileInfo] describing the named file.
		// If there is an error, it will be of type [*PathError].
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			if mode.IsRegular() {
				if mode&0111 != 0 {
					FoundIt = true
					if *minusS == true {
						os.Exit(0)
					}
					if *minusA == true {
						fmt.Println(fullPath)
					} else {
						fmt.Println(fullPath)
						os.Exit(0)
					}
				}
			}
		}
	}

	if FoundIt == false {
		os.Exit(1)
	}
}
