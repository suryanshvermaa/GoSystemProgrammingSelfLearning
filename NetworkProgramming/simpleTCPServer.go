package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("please provide port for server")
		os.Exit(100)
	}
	port := ":" + args[1]
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(100)
	}
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(100)
	}
	defer l.Close()
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(100)
		}
		fmt.Println("->", string(netData))
		c.Write([]byte("server:-> " + netData))
		if strings.TrimSpace(netData) == "STOP" {
			fmt.Println("exiting tcp server")
			return
		}
	}
}
