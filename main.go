package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type ElectricField struct {
	Rect                Rect
	PotentialDifference float64
}

func (e ElectricField) FieldStrength() Vec2 {
	E := e.PotentialDifference / e.Rect.Width() //Uniform Electric Field Strength
	return Vec2{E, 0}
	//points from + to - and I've set that to the useful way of my sim
}

func (e ElectricField) Draw(screen *ebiten.Image) {
	accelRegion := ebiten.NewImage(int(e.Rect.Width()), int(e.Rect.Height()))
	accelRegion.Fill(color.RGBA{250, 50, 50, 100})

	drawOps := ebiten.DrawImageOptions{}
	drawOps.GeoM.Translate(e.Rect.Min.X, e.Rect.Min.Y)

	screen.DrawImage(accelRegion, &drawOps)
}

type Detector struct {
	Rect              Rect
	AcellerationField ElectricField
	ticksElapsed      int
}

func (d *Detector) Update(molecules []*Molecule) {
	d.ticksElapsed++

	for _, m := range molecules {
		if d.Rect.Contains(m.Pos) {
			d.TakeReading(m)
		}
	}
}

func (d *Detector) TakeReading(molecule *Molecule) {
	molecule.Active = false
	z := molecule.Charge //simulate reading the charge
	molecule.Charge = 0
	t := float64(d.ticksElapsed) * DT // in seconds

	E := float64(z) * d.AcellerationField.PotentialDifference // Electrical energy
	v := L / t                                                // Constant velocity
	m := 2 * E / (v * v)                                      // Kinetic energy

	mpz := m / float64(z)

	fmt.Println(molecule.Mass(), mpz)
}

func (d Detector) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(int(d.Rect.Width()), int(d.Rect.Height()))
	img.Fill(color.RGBA{60, 75, 75, 255})

	drawOps := ebiten.DrawImageOptions{}
	drawOps.GeoM.Translate(d.Rect.Min.X, d.Rect.Min.Y)

	screen.DrawImage(img, &drawOps)
}

type Simulation struct {
	accelerationRegion ElectricField
	detector           Detector

	methane         Molecule
	drawableMethane RenderMolecule

	wasIn bool //TEMP just to properly do detector timing for now
}

func NewSimulation() *Simulation {
	s := &Simulation{
		accelerationRegion: ElectricField{
			Rect:                NewRect(100, 150, 300, 750),
			PotentialDifference: 16_000,
		},
		detector: Detector{
			Rect:         NewRect(1400, 150, 1500, 750),
			ticksElapsed: 0,
		},

		methane: Molecule{
			Name:   "methane",
			Active: true,
			Atoms: []struct {
				element *Atom
				count   int
			}{
				{&CARBON, 1},
				{&HYDROGEN, 4},
			},
			Charge: 1,
			Pos:    Vec2{100, 450},
			Vel:    Vec2{0, 0},
		},

		wasIn: true,
	}

	s.drawableMethane = RenderMolecule{
		&s.methane,
		color.RGBA{200, 210, 210, 255},
	}
	s.detector.AcellerationField = s.accelerationRegion

	return s
}

func (s *Simulation) Update() error {

	molecules := []*Molecule{}

	if s.methane.Active {
		s.methane.Update(s.accelerationRegion)
		molecules = append(molecules, &s.methane)
	}

	if !s.accelerationRegion.Rect.Contains(s.methane.Pos) && s.wasIn {
		s.detector.ticksElapsed = 0
		s.wasIn = false
	}

	s.detector.Update(molecules)

	return nil
}

func (s Simulation) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 100, 120, 255})

	s.accelerationRegion.Draw(screen)
	s.detector.Draw(screen)

	s.drawableMethane.Draw(screen)
}

func (s Simulation) Layout(actualWidth, actualHeight int) (screenWidth, screenHeight int) {
	return 1600, 900
}

func main() {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Mass Spectrometry")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	sim := NewSimulation()
	if err := ebiten.RunGame(sim); err != nil {
		panic(err)
	}
}
