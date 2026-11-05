package main

import "fmt"

// ---

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

// ---

type Dress interface {
	getColor() string
}

// ---

type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}

// ---

type CounterTerroristDress struct {
	color string
}

func (c *CounterTerroristDress) getColor() string {
	return c.color
}

// ---

type DressFactory struct {
	dressMap map[string]Dress
}

var dressFactoryInstance = &DressFactory{
	dressMap: make(map[string]Dress),
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.dressMap[dressType] = &TerroristDress{color: "red"}
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = &CounterTerroristDress{color: "green"}
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("Wrong dress type passed")
}

// ---

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := dressFactoryInstance.getDressByType(dressType)
	return &Player{playerType: playerType, dress: dress}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

// ---

type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *Game {
	return &Game{}
}

func (g *Game) addTerrorist(dressType string) {
	g.terrorists = append(g.terrorists, newPlayer("T", dressType))
}

func (g *Game) addCounterTerrorist(dressType string) {
	g.counterTerrorists = append(g.counterTerrorists, newPlayer("CT", dressType))
}

// ---

func main() {
	g := newGame()

	g.addTerrorist(TerroristDressType)
	g.addTerrorist(TerroristDressType)
	g.addTerrorist(TerroristDressType)
	g.addTerrorist(TerroristDressType)

	g.addCounterTerrorist(CounterTerroristDressType)
	g.addCounterTerrorist(CounterTerroristDressType)
	g.addCounterTerrorist(CounterTerroristDressType)

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
