package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const interval = 2 * time.Second

func runCommand(args []string) ([]byte, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("usage: watch <command> [args...]")
	}

	cmd := exec.Command(args[0], args[1:]...)

	return cmd.CombinedOutput()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: watch <command> [args...]")
		os.Exit(1)
	}

	command := os.Args[1:]

	for {
		output, err := runCommand(command)

		// Clear the terminal before printing the next result.
		fmt.Print("\033[H\033[2J")

		fmt.Print(string(output))

		if err != nil {
			fmt.Fprintf(os.Stderr, "\ncommand failed: %v\n", err)
		}

		time.Sleep(interval)
	}
}
