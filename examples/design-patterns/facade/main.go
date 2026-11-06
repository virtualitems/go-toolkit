package main

import "fmt"

type Lights struct{}

func (l *Lights) TurnOff() {
	fmt.Println("Luces apagadas")
}

func (l *Lights) TurnOn() {
	fmt.Println("Luces encendidas")
}

type Projector struct{}

func (p *Projector) TurnOn() {
	fmt.Println("Proyector encendido")
}

func (p *Projector) TurnOff() {
	fmt.Println("Proyector apagado")
}

type SoundSystem struct{}

func (s *SoundSystem) TurnOn() {
	fmt.Println("Sistema de sonido encendido")
}

func (s *SoundSystem) TurnOff() {
	fmt.Println("Sistema de sonido apagado")
}

type HomeTheaterFacade struct {
	lights      *Lights
	projector   *Projector
	soundSystem *SoundSystem
}

func NewHomeTheater() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		lights:      &Lights{},
		projector:   &Projector{},
		soundSystem: &SoundSystem{},
	}
}

func (h *HomeTheaterFacade) WatchMovie() {
	h.lights.TurnOff()
	h.projector.TurnOn()
	h.soundSystem.TurnOn()
	fmt.Println("Película iniciada")
}

func (h *HomeTheaterFacade) EndMovie() {
	h.projector.TurnOff()
	h.soundSystem.TurnOff()
	h.lights.TurnOn()
	fmt.Println("Película terminada")
}

func main() {
	homeTheater := NewHomeTheater()

	homeTheater.WatchMovie()
	fmt.Println()
	homeTheater.EndMovie()
}
