package main

import "fmt"

// ---

type Pizza interface {
	getPrice() int
}

// ---

type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
	return 15
}

// ---

type TomatoTopping struct {
	pizza Pizza
}

func (c *TomatoTopping) getPrice() int {
	return c.pizza.getPrice() + 7
}

// ---

type CheeseTopping struct {
	pizza Pizza
}

func (c *CheeseTopping) getPrice() int {
	return c.pizza.getPrice() + 10
}

// ---

func main() {
	pizza := &VeggieMania{}

	pizzaWithCheese := &CheeseTopping{pizza: pizza}
	pizzaWithCheeseAndTomato := &TomatoTopping{pizza: pizzaWithCheese}

	fmt.Printf("Price of veggieMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
