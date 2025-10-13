package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("10 goroutines\n")
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("goroutine running:", i)
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("exiting...")
}
