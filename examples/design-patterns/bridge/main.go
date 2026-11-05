package main

import "fmt"

// ---

type Computer interface {
	Print()
	SetPrinter(Printer)
}

// ---

type Printer interface {
	PrintFile()
}

// ---

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

// ---

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

// ---

type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

// ---

type HP struct {
}

func (p *HP) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

// ---

func main() {
	hpPrinter := &HP{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()

	winComputer := &Windows{}
	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
}
