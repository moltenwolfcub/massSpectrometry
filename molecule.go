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

type BondedElement struct {
	element  *Element
	children []*BondedElement
}

type BondedAtom struct {
	atom     *Atom
	children []*BondedAtom
}

type Molecule struct {
	Name       string
	Active     bool
	DriftTicks int
	Charge     int
	Pos        Vec2
	Vel        Vec2
	Mass       float64

	ElementalStructure *BondedElement
	AtomicStructure    *BondedAtom
}

func NewMolecule(name string, structure *BondedElement) Molecule {
	m := Molecule{
		Name:               name,
		ElementalStructure: structure,
		Mass:               -1,
	}
	return m
}

func (m *Molecule) SetIsotope() {
	root := m.setIsotopeChildren(m.ElementalStructure)
	m.AtomicStructure = root
}

func (m Molecule) setIsotopeChildren(bondedElement *BondedElement) *BondedAtom {
	atom := &BondedAtom{
		atom: bondedElement.element.GetIsotope(),
	}
	for _, child := range bondedElement.children {
		assignedChild := m.setIsotopeChildren(child)
		atom.children = append(atom.children, assignedChild)
	}
	return atom
}

func (m Molecule) CalcMass() float64 {
	if m.Mass >= 0 {
		return m.Mass
	}

	mass := 0.0
	mass += float64(m.AtomicStructure.atom.AtomicMass)

	toSearch := append(make([]*BondedAtom, 0), m.AtomicStructure.children...)

	for i := 0; i < len(toSearch); i++ {
		e := toSearch[i]
		mass += float64(e.atom.AtomicMass)
		toSearch = append(toSearch, e.children...)
	}

	m.Mass = mass
	return mass
}

func (m *Molecule) Update(electricField ElectricField) {
	var F Vec2

	if electricField.Rect.Contains(m.Pos) {
		F = electricField.FieldStrength().Mul(float64(m.Charge)) //Lorentz Force
	} else {
		m.DriftTicks++
	}

	a := F.Mul(1 / m.CalcMass()) //Newton's 2nd law

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
