package main

import "fmt"

func main() {
	fmt.Println("classic for:")
	for i := 5; i <= 10; i += 1 {
		fmt.Println(i)
	}

	fmt.Println("conditional for:")
	n := 1
	for n <= 3 {
		fmt.Println(n)
		n++
	}

	fmt.Println("infinite for:")
	for {
		fmt.Println("once")
		break
	}

	numbers := []int{10, 20, 30}
	fmt.Println("range with index and value:")
	for index, number := range numbers {
		fmt.Printf("%d: %d\n", index, number)
	}

	fmt.Println("range with index only:")
	for index := range numbers {
		fmt.Println(index)
	}

	fmt.Println("range with value only:")
	for _, number := range numbers {
		fmt.Println(number)
	}

	fmt.Println("range without variables:")
	for range numbers {
		fmt.Println("iteration")
	}

	fmt.Println("range over a string:")
	for index, letter := range "Go!" {
		fmt.Printf("%d: %c\n", index, letter)
	}

	ages := map[string]int{"Ana": 20, "Luis": 25}
	fmt.Println("range over a map:")
	for name, age := range ages {
		fmt.Printf("%s: %d\n", name, age)
	}

	values := make(chan int, 2)
	values <- 1
	values <- 2
	close(values)
	fmt.Println("range over a channel:")
	for value := range values {
		fmt.Println(value)
	}
}
