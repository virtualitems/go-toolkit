package main

import (
	"fmt"
	"log"
)

// ---

type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

// ---

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	currentState  State
	itemCount     int
	itemPrice     int
}

func newVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{itemCount: itemCount, itemPrice: itemPrice}
	v.hasItem = &HasItemState{vendingMachine: v}
	v.itemRequested = &ItemRequestedState{vendingMachine: v}
	v.hasMoney = &HasMoneyState{vendingMachine: v}
	v.noItem = &NoItemState{vendingMachine: v}
	v.setState(v.hasItem)
	return v
}

func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount += count
}

// ---

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

// ---

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (i *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}
func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}
func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Println("Item requested")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

// ---

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}
func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

// ---

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount--
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

// ---

func main() {
	vendingMachine := newVendingMachine(1, 10)

	if err := vendingMachine.requestItem(); err != nil {
		log.Fatalf(err.Error())
	}
	if err := vendingMachine.insertMoney(10); err != nil {
		log.Fatalf(err.Error())
	}
	if err := vendingMachine.dispenseItem(); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	if err := vendingMachine.addItem(2); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	if err := vendingMachine.requestItem(); err != nil {
		log.Fatalf(err.Error())
	}
	if err := vendingMachine.insertMoney(10); err != nil {
		log.Fatalf(err.Error())
	}
	if err := vendingMachine.dispenseItem(); err != nil {
		log.Fatalf(err.Error())
	}
}
