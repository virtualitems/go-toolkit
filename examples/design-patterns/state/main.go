package main

import "fmt"

type Event string

const TimerExpired Event = "terminó el tiempo"

type State interface {
	Name() string
	Handle(event Event, light *TrafficLight)
}

type TrafficLight struct {
	state State
}

func NewTrafficLight() *TrafficLight {
	return &TrafficLight{state: &RedState{}}
}

func (t *TrafficLight) Trigger(event Event) {
	fmt.Printf("Estado actual: %s\n", t.state.Name())
	fmt.Printf("Evento: %s\n", event)
	t.state.Handle(event, t)
}

func (t *TrafficLight) transitionTo(next State) {
	fmt.Printf("Transición: %s -> %s\n\n", t.state.Name(), next.Name())
	t.state = next
}

type RedState struct{}

func (r *RedState) Name() string {
	return "rojo"
}

func (r *RedState) Handle(event Event, light *TrafficLight) {
	if event == TimerExpired {
		light.transitionTo(&GreenState{})
	}
}

type GreenState struct{}

func (g *GreenState) Name() string {
	return "verde"
}

func (g *GreenState) Handle(event Event, light *TrafficLight) {
	if event == TimerExpired {
		light.transitionTo(&YellowState{})
	}
}

type YellowState struct{}

func (y *YellowState) Name() string {
	return "amarillo"
}

func (y *YellowState) Handle(event Event, light *TrafficLight) {
	if event == TimerExpired {
		light.transitionTo(&RedState{})
	}
}

func main() {
	light := NewTrafficLight()

	light.Trigger(TimerExpired)
	light.Trigger(TimerExpired)
	light.Trigger(TimerExpired)
}
