package main

import "fmt"

// ---

type Command interface {
	execute()
}

// ---

type Device interface {
	on()
	off()
}

// ---

type TV struct {
	isRunning bool
}

func (t *TV) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *TV) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

// ---

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

// ---

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// ---

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

// ---

func main() {
	tv := &TV{}

	onButton := &Button{command: &OnCommand{device: tv}}
	onButton.press()

	offButton := &Button{command: &OffCommand{device: tv}}
	offButton.press()
}
