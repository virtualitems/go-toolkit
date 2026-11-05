package main

import "fmt"

type SportswearFactory interface {
	makeShoe() Shoe
	makeShirt() Shirt
}

// ---

type Shoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type ShoeDetails struct {
	logo string
	size int
}

func (s *ShoeDetails) setLogo(logo string) {
	s.logo = logo
}

func (s *ShoeDetails) getLogo() string {
	return s.logo
}

func (s *ShoeDetails) setSize(size int) {
	s.size = size
}

func (s *ShoeDetails) getSize() int {
	return s.size
}

type AdidasShoe struct {
	ShoeDetails
}

type NikeShoe struct {
	ShoeDetails
}

// ---

type Shirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type ShirtDetails struct {
	logo string
	size int
}

func (s *ShirtDetails) setLogo(logo string) {
	s.logo = logo
}

func (s *ShirtDetails) getLogo() string {
	return s.logo
}

func (s *ShirtDetails) setSize(size int) {
	s.size = size
}

func (s *ShirtDetails) getSize() int {
	return s.size
}

type AdidasShirt struct {
	ShirtDetails
}

type NikeShirt struct {
	ShirtDetails
}

// ---

type Adidas struct {
}

func (Adidas) makeShoe() Shoe {
	return &AdidasShoe{ShoeDetails: ShoeDetails{logo: "adidas", size: 14}}
}

func (Adidas) makeShirt() Shirt {
	return &AdidasShirt{ShirtDetails: ShirtDetails{logo: "adidas", size: 14}}
}

// ---

type Nike struct {
}

func (Nike) makeShoe() Shoe {
	return &NikeShoe{ShoeDetails: ShoeDetails{logo: "nike", size: 14}}
}

func (Nike) makeShirt() Shirt {
	return &NikeShirt{ShirtDetails: ShirtDetails{logo: "nike", size: 14}}
}

// ---

type SportswearAbstractFactory struct {
}

func (SportswearAbstractFactory) Adidas() SportswearFactory {
	return Adidas{}
}

func (SportswearAbstractFactory) Nike() SportswearFactory {
	return Nike{}
}

// ---

func printShoeDetails(s Shoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func printShirtDetails(s Shirt) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func main() {
	factory := SportswearAbstractFactory{}
	adidasFactory := factory.Adidas()
	nikeFactory := factory.Nike()

	printShoeDetails(adidasFactory.makeShoe())
	printShirtDetails(adidasFactory.makeShirt())

	printShoeDetails(nikeFactory.makeShoe())
	printShirtDetails(nikeFactory.makeShirt())
}
