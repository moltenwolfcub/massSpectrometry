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
	E := Metre(e.PotentialDifference) / e.Rect.Width() //Uniform Electric Field Strength
	return Vec2{E, 0}
	//points from + to - and I've set that to the useful way of my sim
}

func (e ElectricField) Draw(screen *ebiten.Image) {
	accelRegion := ebiten.NewImage(int(e.Rect.Width().ToPixel()), int(e.Rect.Height().ToPixel()))
	accelRegion.Fill(color.RGBA{250, 50, 50, 100})

	drawOps := ebiten.DrawImageOptions{}
	drawOps.GeoM.Translate(float64(e.Rect.Min.X.ToPixel()), float64(e.Rect.Min.Y.ToPixel()))

	screen.DrawImage(accelRegion, &drawOps)
}

type Detector struct {
	Rect              Rect
	AcellerationField ElectricField
	ticksElapsed      Tick
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
	t := d.ticksElapsed.ToSecond()

	E := float64(z) * d.AcellerationField.PotentialDifference // Electrical energy
	v := L / Metre(t)                                         // Constant velocity
	m := 2 * Metre(E) / (v * v)                               // Kinetic energy

	mpz := float64(m) / float64(z)

	fmt.Println(molecule.Mass(), mpz)
}

func (d Detector) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(int(d.Rect.Width().ToPixel()), int(d.Rect.Height().ToPixel()))
	img.Fill(color.RGBA{60, 75, 75, 255})

	drawOps := ebiten.DrawImageOptions{}
	drawOps.GeoM.Translate(float64(d.Rect.Min.X.ToPixel()), float64(d.Rect.Min.Y.ToPixel()))

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
			Rect: NewRect(
				Pixel(100).ToMetre(), Pixel(150).ToMetre(),
				Pixel(300).ToMetre(), Pixel(750).ToMetre(),
			),
			PotentialDifference: 16_000,
		},
		detector: Detector{
			Rect: NewRect(
				Pixel(1400).ToMetre(), Pixel(150).ToMetre(),
				Pixel(1500).ToMetre(), Pixel(750).ToMetre(),
			),
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
			Pos:    Vec2{Pixel(100).ToMetre(), Pixel(450).ToMetre()},
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
