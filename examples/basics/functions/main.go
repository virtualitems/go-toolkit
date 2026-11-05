package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(n1 int, n2 int) int {
	return n1 + n2
}

func addx(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	args := os.Args[1:]
	argsLen := len(args)

	if argsLen < 2 {
		fmt.Fprintln(os.Stderr, "Error: there must be at least two numbers as arguments")
		os.Exit(1)
	}

	numbers := make([]int, argsLen)

	for i, arg := range args {
		number, err := strconv.Atoi(arg)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %q is not a number\n", arg)
			os.Exit(1)
		}

		numbers[i] = number
	}

	if argsLen == 2 {
		result := add(numbers[0], numbers[1])
		fmt.Println("result:", result)
		return
	}

	result := addx(numbers...)
	fmt.Println("result:", result)
}
