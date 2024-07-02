package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	// Define flags
	command := flag.String("command", "", "Specify the command (echo, ls, execute)")
	message := flag.String("message", "", "Message to echo")
	path := flag.String("path", "", "File path to execute")

	// Parse flags
	flag.Parse()
	// Check if required flags are provided
	if *command == "" {
		fmt.Println("Usage: go run main.go --command <command>")
		fmt.Println("Available commands: echo, ls, execute")
		return
	}
	switch *command {
	case "echo":
		if *message == "" {
			fmt.Println("Message is required for 'echo' command")
			fmt.Println("Usage: go run main.go --command echo --message hello")
			return
		}
		fmt.Println(*message)
	case "ls":
		lsCmd := exec.Command("ls")
		lsCmd.Stdout = os.Stdout
		lsCmd.Stderr = os.Stderr
		if err := lsCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error running ls command: %v\n", err)
			os.Exit(1)
		}
	case "execute":
		if *path == "" {
			fmt.Println("Path is required for 'execute' command")
			fmt.Println("Usage: go run main.go --command execute --path code.sh")
			return
		}
		cmd := exec.Command("sh", *path)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error running shell script: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", *command)
		fmt.Println("Available commands: echo, ls, execute")
	}
}
