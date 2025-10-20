package main

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Element struct {
	Name     string
	Isotopes []struct {
		atom      *Atom
		abundance float64
	}
}

func (e Element) GetIsotope() *Atom {
	r := rand.Float64()

	for _, i := range e.Isotopes {
		r -= i.abundance
		if r <= 0 {
			return i.atom
		}
	}

	return e.Isotopes[len(e.Isotopes)-1].atom
}

type Atom struct {
	Name         string
	AtomicNumber int
	AtomicMass   int
}

type Molecule struct {
	Name       string
	Active     bool
	DriftTicks int
	Elements   []struct {
		element *Element
		count   int
	}
	Atoms  []*Atom
	Charge int
	Pos    Vec2
	Vel    Vec2
}

func (m *Molecule) SetIsotope() {
	m.Atoms = make([]*Atom, 0)
	for _, e := range m.Elements {
		for i := 0; i < e.count; i++ {
			isotope := e.element.GetIsotope()
			m.Atoms = append(m.Atoms, isotope)
		}
	}
}

func (m Molecule) Mass() float64 {
	mass := 0.0
	for _, a := range m.Atoms {
		elementMass := a.AtomicMass
		mass += float64(elementMass)
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
