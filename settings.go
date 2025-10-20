package main

import "image/color"

const (
	DT   float64 = 0.001 //Seconds per tick
	PXPM int     = 100   //pixels per metre
)

// atoms
var (
	A_HYDROGEN = Atom{
		Name:         "hydrogen",
		AtomicNumber: 1,
		AtomicMass:   1,
	}
	A_DEUTERIUM = Atom{
		Name:         "deuterium",
		AtomicNumber: 1,
		AtomicMass:   2,
	}
	A_CARBON_12 = Atom{
		Name:         "carbon-12",
		AtomicNumber: 6,
		AtomicMass:   12,
	}
	A_CARBON_13 = Atom{
		Name:         "carbon-13",
		AtomicNumber: 6,
		AtomicMass:   13,
	}
)

// elements
var (
	HYDROGEN = Element{
		Name: "hydrogen",
		Isotopes: []struct {
			atom      *Atom
			abundance float64
		}{
			{
				atom:      &A_HYDROGEN,
				abundance: 0.9855,
			},
			{
				atom:      &A_DEUTERIUM,
				abundance: 0.0145,
			},
		},
	}
	CARBON = Element{
		Name: "carbon",
		Isotopes: []struct {
			atom      *Atom
			abundance float64
		}{
			{
				atom:      &A_CARBON_12,
				abundance: 0.9884,
			},
			{
				atom:      &A_CARBON_13,
				abundance: 0.0116,
			},
		},
	}

//	CARBON = Atom{
//		Name:         "carbon",
//		AtomicNumber: 6,
//		AtomicMass:   12,
//	}
//
//	COPPER63 = Atom{
//		Name:         "copper 63",
//		AtomicNumber: 29,
//		AtomicMass:   63,
//	}
//
//	COPPER65 = Atom{
//		Name:         "copper 65",
//		AtomicNumber: 29,
//		AtomicMass:   65,
//	}
//
//	OXYGEN = Atom{
//		Name:         "oxygen",
//		AtomicNumber: 8,
//		AtomicMass:   16,
//	}
//
//	NITROGEN = Atom{
//		Name:         "nitrogen",
//		AtomicNumber: 7,
//		AtomicMass:   14,
//	}
)

var MOLECULES []struct {
	m Molecule
	c color.Color
} = []struct {
	m Molecule
	c color.Color
}{
	{
		m: Molecule{
			Name: "methane",
			Elements: []struct {
				element *Element
				count   int
			}{{&CARBON, 1}, {&HYDROGEN, 4}},
		},
		c: color.RGBA{250, 50, 50, 255},
	},
	// {
	// 	m: Molecule{
	// 		Name: "ethanol",
	// 		Atoms: []struct {
	// 			element *Atom
	// 			count   int
	// 		}{{&CARBON, 2}, {&HYDROGEN, 5}, {&OXYGEN, 1}},
	// 	},
	// 	c: color.RGBA{50, 150, 50, 255},
	// },
	// {
	// 	m: Molecule{
	// 		Name: "propane",
	// 		Atoms: []struct {
	// 			element *Atom
	// 			count   int
	// 		}{{&CARBON, 3}, {&HYDROGEN, 8}},
	// 	},
	// 	c: color.RGBA{200, 75, 50, 255},
	// },
	// {
	// 	m: Molecule{
	// 		Name: "copper 63",
	// 		Atoms: []struct {
	// 			element *Atom
	// 			count   int
	// 		}{{&COPPER63, 1}},
	// 	},
	// 	c: color.RGBA{114, 73, 12, 255},
	// },
	// {
	// 	m: Molecule{
	// 		Name: "copper 65",
	// 		Atoms: []struct {
	// 			element *Atom
	// 			count   int
	// 		}{{&COPPER65, 1}},
	// 	},
	// 	c: color.RGBA{114, 73, 12, 255},
	// },
	// {
	// 	m: Molecule{
	// 		Name: "water",
	// 		Atoms: []struct {
	// 			element *Atom
	// 			count   int
	// 		}{{&HYDROGEN, 2}, {&OXYGEN, 1}},
	// 	},
	// 	c: color.RGBA{45, 45, 220, 255},
	// },
	// {
	// 	m: Molecule{
	// 		Name: "tnt",
	// 		Atoms: []struct {
	// 			element *Atom
	// 			count   int
	// 		}{{&CARBON, 7}, {&HYDROGEN, 8}, {&NITROGEN, 3}, {&OXYGEN, 6}},
	// 	},
	// 	c: color.RGBA{50, 50, 50, 255},
	// },
	// {
	// 	m: Molecule{
	// 		Name: "Hydrogen",
	// 		Atoms: []struct {
	// 			element *Atom
	// 			count   int
	// 		}{{&HYDROGEN, 2}},
	// 	},
	// 	c: color.RGBA{170, 170, 50, 255},
	// },
}
