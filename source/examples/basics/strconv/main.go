package main

import (
    "fmt"
    "os"
	"strconv"
)

func main() {
    args := os.Args[1:]

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "no arguments")
		os.Exit(1)
	}

	number, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Fprintf(os.Stderr, "First argument must be a number")
		os.Exit(1)
	}

	fmt.Println("Number:", number)
}
