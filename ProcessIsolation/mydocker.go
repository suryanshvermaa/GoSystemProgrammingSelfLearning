package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mydocker run <command>")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "run":
		run()
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}

func run() {
	fmt.Printf("Running %v in new namespaces...\n", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// namespace isolation
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWIPC,
	}
	syscall.Sethostname([]byte("suryansh-container"))
	fmt.Printf("Parent PID: %d\n", os.Getpid())

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}
