package main

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
	AtomicMass   float32
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
