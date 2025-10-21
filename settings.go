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

	A_LITHIUM_6 = Atom{
		Name:         "lithium-6",
		AtomicNumber: 3,
		AtomicMass:   6,
	}
	A_LITHIUM_7 = Atom{
		Name:         "lithium-7",
		AtomicNumber: 3,
		AtomicMass:   7,
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
		Name:         "nitrogen-14",
		AtomicNumber: 7,
		AtomicMass:   14,
	}
	A_NITROGEN_15 = Atom{
		Name:         "nitrogen-15",
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

	A_NEON_20 = Atom{
		Name:         "neon-20",
		AtomicNumber: 10,
		AtomicMass:   20,
	}
	A_NEON_21 = Atom{
		Name:         "neon-21",
		AtomicNumber: 10,
		AtomicMass:   21,
	}
	A_NEON_22 = Atom{
		Name:         "neon-22",
		AtomicNumber: 10,
		AtomicMass:   22,
	}

	A_SODIUM_23 = Atom{
		Name:         "neon-22",
		AtomicNumber: 11,
		AtomicMass:   23,
	}

	A_SULFUR_32 = Atom{
		Name:         "sulfur-32",
		AtomicNumber: 16,
		AtomicMass:   32,
	}
	A_SULFUR_33 = Atom{
		Name:         "sulfur-33",
		AtomicNumber: 16,
		AtomicMass:   33,
	}
	A_SULFUR_34 = Atom{
		Name:         "sulfur-34",
		AtomicNumber: 16,
		AtomicMass:   34,
	}
	A_SULFUR_36 = Atom{
		Name:         "sulfur-36",
		AtomicNumber: 16,
		AtomicMass:   36,
	}

	A_CHLORINE_35 = Atom{
		Name:         "chlorine-35",
		AtomicNumber: 17,
		AtomicMass:   35,
	}
	A_CHLORINE_37 = Atom{
		Name:         "chlorine-35",
		AtomicNumber: 17,
		AtomicMass:   37,
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
	LITHIUM = Element{
		Name: "lithium",
		Isotopes: []struct {
			atom      *Atom
			abundance float64
		}{
			{
				atom:      &A_LITHIUM_6,
				abundance: 0.078,
			},
			{
				atom:      &A_LITHIUM_7,
				abundance: 0.922,
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
	NEON = Element{
		Name: "neon",
		Isotopes: []struct {
			atom      *Atom
			abundance float64
		}{
			{
				atom:      &A_NEON_20,
				abundance: 0.9048,
			},
			{
				atom:      &A_NEON_21,
				abundance: 0.0027,
			},
			{
				atom:      &A_NEON_22,
				abundance: 0.0925,
			},
		},
	}
	SODIUM = Element{
		Name: "sodium",
		Isotopes: []struct {
			atom      *Atom
			abundance float64
		}{
			{
				atom:      &A_SODIUM_23,
				abundance: 1,
			},
		},
	}
	CHLORINE = Element{
		Name: "chlorine",
		Isotopes: []struct {
			atom      *Atom
			abundance float64
		}{
			{
				atom:      &A_CHLORINE_35,
				abundance: 0.758,
			},
			{
				atom:      &A_CHLORINE_37,
				abundance: 0.242,
			},
		},
	}
	SULFUR = Element{
		Name: "sulfur",
		Isotopes: []struct {
			atom      *Atom
			abundance float64
		}{
			{
				atom:      &A_SULFUR_32,
				abundance: 0.948562,
			},
			{
				atom:      &A_SULFUR_33,
				abundance: 0.00763,
			},
			{
				atom:      &A_SULFUR_34,
				abundance: 0.04365,
			},
			{
				atom:      &A_SULFUR_36,
				abundance: 0.000158,
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
			Name: "hydrogen",
			Elements: []struct {
				element *Element
				count   int
			}{{&HYDROGEN, 2}},
		},
		c: color.RGBA{170, 170, 50, 255},
	},
	{
		m: Molecule{
			Name: "neon",
			Elements: []struct {
				element *Element
				count   int
			}{{&NEON, 1}},
		},
		c: color.RGBA{100, 8, 150, 255},
	},
	{
		m: Molecule{
			Name: "sodium",
			Elements: []struct {
				element *Element
				count   int
			}{{&SODIUM, 1}},
		},
		c: color.RGBA{100, 150, 0, 255},
	},
	{
		m: Molecule{
			Name: "chlorine",
			Elements: []struct {
				element *Element
				count   int
			}{{&CHLORINE, 2}},
		},
		c: color.RGBA{50, 75, 0, 255},
	},
}
