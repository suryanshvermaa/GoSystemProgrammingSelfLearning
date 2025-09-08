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
	err := os.Remove(file)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

}
