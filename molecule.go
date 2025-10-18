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
		AtomicMass:   12,
	}

	HYDROGEN = Atom{
		Name:         "hydrogen",
		AtomicNumber: 1,
		AtomicMass:   1,
	}
)

type Atom struct {
	Name         string
	AtomicNumber int
	AtomicMass   int
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

func (m Molecule) Mass() float64 {
	mass := 0.0
	for _, a := range m.Atoms {
		elementMass := a.element.AtomicMass
		mass += float64(elementMass * a.count)
	}
	return mass
}

func (m *Molecule) Update(electricField ElectricField) {
	var F Vec2

	if electricField.Rect.Contains(m.Pos) {
		F = electricField.FieldStrength().Mul(float64(m.Charge)) //Lorentz Force
	}

	a := F.Mul(1 / m.Mass()) //Newton's 2nd law

	u := m.Vel
	v := u.Add(a.Mul(DT))                        //SUVAT 2
	ds := u.Mul(DT).Add(a.Mul(DT * DT).Mul(0.5)) //SUVAT 3

	m.Pos = m.Pos.Add(ds)
	m.Vel = v
}

type RenderMolecule struct {
	Molecule *Molecule
	Color    color.Color
}

func (m RenderMolecule) Draw(screen *ebiten.Image) {
	var size float32 = 10
	vector.FillCircle(screen, float32(m.Molecule.Pos.X), float32(m.Molecule.Pos.Y), size, m.Color, true)
}
