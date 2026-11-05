package main

import "fmt"

// ---

type EuropeanOutlet interface {
	PlugIntoEuropeanOutlet()
}

// ---

type Traveler struct {
}

func (t *Traveler) ChargeDevice(outlet EuropeanOutlet) {
	fmt.Println("Traveler looks for a place to charge.")
	outlet.PlugIntoEuropeanOutlet()
}

// ---

type EuropeanLaptop struct {
}

func (e *EuropeanLaptop) PlugIntoEuropeanOutlet() {
	fmt.Println("European laptop plugs in directly.")
}

// ---

type AmericanPhone struct {
}

func (a *AmericanPhone) PlugIntoAmericanOutlet() {
	fmt.Println("American phone plugs into flat-pin outlet.")
}

// ---

type TravelAdapter struct {
	phone *AmericanPhone
}

func (t *TravelAdapter) PlugIntoEuropeanOutlet() {
	fmt.Println("Travel adapter converts european outlet to american.")
	t.phone.PlugIntoAmericanOutlet()
}

// ---

func main() {
	traveler := &Traveler{}

	laptop := &EuropeanLaptop{}
	traveler.ChargeDevice(laptop)

	phone := &AmericanPhone{}
	adapter := &TravelAdapter{phone: phone}
	traveler.ChargeDevice(adapter)
}
