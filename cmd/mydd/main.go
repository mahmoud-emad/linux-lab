package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func usage() string {
	return "Usage: mydd if=<source_file> of=<destination>"
}

func argumentCheck(argument string, keyword string) (string, error) {
	parts := strings.SplitN(argument, "=", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf(usage())
	}
	if strings.ToLower(parts[0]) != keyword {
		return "", fmt.Errorf(usage())
	}
	return parts[1], nil
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage())
		return
	}

	sourceFileValue := args[1]
	fvalue, err := argumentCheck(sourceFileValue, "if")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	destFileValue := args[2]
	dvalue, err := argumentCheck(destFileValue, "of")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Create (or truncate) the destination file.
	destFile, err := os.Create(dvalue)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Read the first file, make sure it is exists.
	defer destFile.Close()
	sourceFile, err := os.Open(fvalue)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer sourceFile.Close()
	data := make([]byte, 1<<20) // Can be changed with the block size, bytes size

	for {
		bytesRead, err := sourceFile.Read(data)

		if bytesRead > 0 {
			_, err = destFile.Write(data[:bytesRead])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

	}
}
