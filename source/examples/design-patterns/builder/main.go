package main

import "fmt"

// ---

type HouseBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

// ---

type House struct {
	windowType string
	doorType   string
	floor      int
}

// ---

type StandardHouseBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *StandardHouseBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *StandardHouseBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *StandardHouseBuilder) setNumFloor() {
	b.floor = 2
}

func (b *StandardHouseBuilder) getHouse() House {
	return House{doorType: b.doorType, windowType: b.windowType, floor: b.floor}
}

// ---

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{doorType: b.doorType, windowType: b.windowType, floor: b.floor}
}

// ---

type HouseDirector struct {
	builder HouseBuilder
}

func (d *HouseDirector) setBuilder(b HouseBuilder) {
	d.builder = b
}

func (d *HouseDirector) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

// ---

func main() {
	director := &HouseDirector{builder: &StandardHouseBuilder{}}
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(&IglooBuilder{})
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}
