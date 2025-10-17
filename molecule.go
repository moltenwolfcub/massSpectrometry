package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	CARBON = Atom{
		Name:         "carbon",
		AtomicNumber: 6,
		AtomicMass:   12.0,
	}

	HYDROGEN = Atom{
		Name:         "hydrogen",
		AtomicNumber: 1,
		AtomicMass:   1.0,
	}
)

type Atom struct {
	Name         string
	AtomicNumber int
	AtomicMass   float32 //should probably be int for a given atom
}

type Molecule struct {
	Name  string
	Atoms []struct {
		element *Atom
		count   int
	}
	Charge int
	Pos    Vec2
	Vel    Vec2
}

type RenderMolecule struct {
	Molecule *Molecule
	Color    color.Color
}

func (m RenderMolecule) Draw(screen *ebiten.Image) {
	var size float32 = 10
	vector.FillCircle(screen, m.Molecule.Pos.X, m.Molecule.Pos.Y, size, m.Color, false)
}
