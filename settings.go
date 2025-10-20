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

	A_NITROGEN_14 = Atom{
		Name:         "nitrogen",
		AtomicNumber: 7,
		AtomicMass:   14,
	}
	A_NITROGEN_15 = Atom{
		Name:         "nitrogen",
		AtomicNumber: 7,
		AtomicMass:   15,
	}

	A_OXYGEN_16 = Atom{
		Name:         "oxygen-16",
		AtomicNumber: 8,
		AtomicMass:   16,
	}
	A_OXYGEN_17 = Atom{
		Name:         "oxygen-17",
		AtomicNumber: 8,
		AtomicMass:   17,
	}
	A_OXYGEN_18 = Atom{
		Name:         "oxygen-18",
		AtomicNumber: 8,
		AtomicMass:   18,
	}

	A_COPPER_63 = Atom{
		Name:         "copper-63",
		AtomicNumber: 29,
		AtomicMass:   63,
	}
	A_COPPER_65 = Atom{
		Name:         "copper-65",
		AtomicNumber: 29,
		AtomicMass:   65,
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
	NITROGEN = Element{
		Name: "nitrogen",
		Isotopes: []struct {
			atom      *Atom
			abundance float64
		}{
			{
				atom:      &A_NITROGEN_14,
				abundance: 0.99578,
			},
			{
				atom:      &A_NITROGEN_15,
				abundance: 0.00422,
			},
		},
	}
	OXYGEN = Element{
		Name: "oxygen",
		Isotopes: []struct {
			atom      *Atom
			abundance float64
		}{
			{
				atom:      &A_OXYGEN_16,
				abundance: 0.99738,
			},
			{
				atom:      &A_OXYGEN_17,
				abundance: 0.00040,
			},
			{
				atom:      &A_OXYGEN_18,
				abundance: 0.00222,
			},
		},
	}
	COPPER = Element{
		Name: "copper",
		Isotopes: []struct {
			atom      *Atom
			abundance float64
		}{
			{
				atom:      &A_COPPER_63,
				abundance: 0.6915,
			},
			{
				atom:      &A_COPPER_65,
				abundance: 0.3085,
			},
		},
	}
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
	{
		m: Molecule{
			Name: "copper",
			Elements: []struct {
				element *Element
				count   int
			}{{&COPPER, 1}},
		},
		c: color.RGBA{114, 73, 12, 255},
	},
	{
		m: Molecule{
			Name: "ethanol",
			Elements: []struct {
				element *Element
				count   int
			}{{&CARBON, 2}, {&HYDROGEN, 6}, {&OXYGEN, 1}},
		},
		c: color.RGBA{50, 150, 50, 255},
	},
	{
		m: Molecule{
			Name: "propane",
			Elements: []struct {
				element *Element
				count   int
			}{{&CARBON, 3}, {&HYDROGEN, 8}},
		},
		c: color.RGBA{200, 75, 50, 255},
	},
	{
		m: Molecule{
			Name: "water",
			Elements: []struct {
				element *Element
				count   int
			}{{&HYDROGEN, 2}, {&OXYGEN, 1}},
		},
		c: color.RGBA{45, 45, 220, 255},
	},
	{
		m: Molecule{
			Name: "tnt",
			Elements: []struct {
				element *Element
				count   int
			}{{&CARBON, 7}, {&HYDROGEN, 5}, {&NITROGEN, 3}, {&OXYGEN, 6}},
		},
		c: color.RGBA{50, 50, 50, 255},
	},
	{
		m: Molecule{
			Name: "Hydrogen",
			Elements: []struct {
				element *Element
				count   int
			}{{&HYDROGEN, 2}},
		},
		c: color.RGBA{170, 170, 50, 255},
	},
}
