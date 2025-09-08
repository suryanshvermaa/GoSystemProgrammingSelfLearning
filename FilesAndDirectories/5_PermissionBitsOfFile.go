package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("please provide arguments")
		os.Exit(1)
	}

	file := arguments[1]
	info, err := os.Stat(file)

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	mode := info.Mode()
	fmt.Println(file, ": ", mode)
}
