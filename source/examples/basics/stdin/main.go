package main

import (
	"fmt"
	"os"
)

func main() {
	var args []string = os.Args[1:]

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "no arguments")
		os.Exit(1)
	}

	first := args[0]
	last := args[len(args) - 1]

	fmt.Println("All args:", os.Args);
	fmt.Println("All args without executable:", args);
	fmt.Println("First arg:", first);
	fmt.Println("Last arg:", last);
}
