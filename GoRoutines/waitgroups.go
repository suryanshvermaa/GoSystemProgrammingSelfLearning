package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("waiting groups")
	var waitGroup sync.WaitGroup
	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer waitGroup.Done()
			fmt.Printf("%dth goroutine running\n", i)
		}(i)
	}
	waitGroup.Wait()
	println("exiting...")
}
