package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(path string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = io.Copy(os.Stdout, file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Error: there must be one path argument")
		os.Exit(1)
	}

	readFile(args[0])
}