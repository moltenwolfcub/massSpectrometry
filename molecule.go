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
	COPPER63 = Atom{
		Name:         "copper 63",
		AtomicNumber: 29,
		AtomicMass:   63,
	}
	COPPER65 = Atom{
		Name:         "copper 65",
		AtomicNumber: 29,
		AtomicMass:   65,
	}
	OXYGEN = Atom{
		Name:         "oxygen",
		AtomicNumber: 8,
		AtomicMass:   16,
	}
)

type Atom struct {
	Name         string
	AtomicNumber int
	AtomicMass   int
}

type Molecule struct {
	Name       string
	Active     bool
	DriftTicks int
	Atoms      []struct {
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
	} else {
		m.DriftTicks++
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
	vector.FillCircle(screen, float32(m.Molecule.Pos.X)*float32(PXPM), float32(m.Molecule.Pos.Y)*float32(PXPM), size, m.Color, true)
}
