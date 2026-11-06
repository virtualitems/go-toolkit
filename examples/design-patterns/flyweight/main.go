package main

import "fmt"

type Coffee struct {
	name  string
	price int
}

type Order struct {
	customer string
	table    int
	coffee   *Coffee
}

func (o *Order) print() {
	fmt.Printf(
		"%s pidió %s en la mesa %d por $%d\n",
		o.customer,
		o.coffee.name,
		o.table,
		o.coffee.price,
	)
}

func main() {
	coffeeWithMilk := &Coffee{name: "Café con leche", price: 5000}
	chocolate := &Coffee{name: "Chocolate", price: 6000}

	orders := []*Order{
		{customer: "Ana", table: 1, coffee: coffeeWithMilk},
		{customer: "Luis", table: 2, coffee: coffeeWithMilk},
		{customer: "Marta", table: 3, coffee: chocolate},
	}

	for _, order := range orders {
		order.print()
	}

	fmt.Printf("\nPedidos creados: %d\n", len(orders))
	fmt.Println("Cafés creados: 2")
	fmt.Printf("Ana y Luis comparten el mismo objeto Coffee: %t\n", orders[0].coffee == orders[1].coffee)
}
