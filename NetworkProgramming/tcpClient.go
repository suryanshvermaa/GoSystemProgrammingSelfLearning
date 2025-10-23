package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("please provide host:port")
		os.Exit(100)
	}
	connect := arguments[1]
	c, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(100)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Println(message)
		if strings.TrimSpace(text) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
