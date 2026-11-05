package main

import (
	"fmt"
)

// ---

type Furniture interface {
	setName(name string)
	setPrice(price float64)
	getName() string
	getPrice() float64
}

// ---

type FurnitureDetails struct {
	name  string
	price float64
}

func (p *FurnitureDetails) setName(name string) {
	p.name = name
}

func (p *FurnitureDetails) getName() string {
	return p.name
}

func (p *FurnitureDetails) setPrice(price float64) {
	p.price = price
}

func (p *FurnitureDetails) getPrice() float64 {
	return p.price
}

// ---

type Chair struct {
	FurnitureDetails
}

type Table struct {
	FurnitureDetails
}

// ---

type FurnitureFactory struct {
}

func (FurnitureFactory) Chair() Furniture {
	return &Chair{FurnitureDetails: FurnitureDetails{name: "Chair", price: 49.99}}
}

func (FurnitureFactory) Table() Furniture {
	return &Table{FurnitureDetails: FurnitureDetails{name: "Table", price: 149.99}}
}

// ---

func printDetails(p Furniture) {
	fmt.Printf("Product: %s\n", p.getName())
	fmt.Printf("Price: %.2f\n", p.getPrice())
}

func main() {
	factory := FurnitureFactory{}
	chair := factory.Chair()
	table := factory.Table()

	printDetails(chair)
	printDetails(table)
}
