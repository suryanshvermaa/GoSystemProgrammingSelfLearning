package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	minusO := flag.Bool("o", false, "o")
	minusC := flag.Bool("c", false, "c")
	minusK := flag.Int("k", 0, "k")
	flag.Parse()

	fmt.Println("-o", *minusO)
	fmt.Println("-c", *minusC)
	fmt.Println("-k", *minusK)

	for i, val := range os.Args {
		fmt.Println(i, ":", val)
	}
}
