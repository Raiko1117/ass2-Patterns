package main

import "fmt"

// IGun is the interface for all guns.
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Gun is a concrete implementation of the IGun interface.
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

// Ak47 is a specific gun type.
type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47",
			power: 4,
		},
	}
}

type P90 struct {
	Gun
}

func newP90() IGun {
	return &P90{
		Gun: Gun{
			name:  "P90",
			power: 1,
		},
	}
}

// Factory function for creating guns.
func createGun(gunType string) (IGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "P90":
		return newP90(), nil
	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}
