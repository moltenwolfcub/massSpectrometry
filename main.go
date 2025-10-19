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
	accelRegion := ebiten.NewImage(int(e.Rect.Width()*float64(PXPM)), int(e.Rect.Height()*float64(PXPM)))
	accelRegion.Fill(color.RGBA{250, 50, 50, 100})

	drawOps := ebiten.DrawImageOptions{}
	drawOps.GeoM.Translate(float64(e.Rect.Min.X*float64(PXPM)), float64(e.Rect.Min.Y*float64(PXPM)))

	screen.DrawImage(accelRegion, &drawOps)
}

type Detector struct {
	Rect              Rect
	AcellerationField ElectricField
}

func (d *Detector) Update(molecules []*Molecule) {
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
	t := float64(molecule.DriftTicks) * DT

	L := d.Rect.Min.X - d.AcellerationField.Rect.Max.X

	E := float64(z) * d.AcellerationField.PotentialDifference // Electrical energy
	v := L / t                                                // Constant velocity
	m := 2 * E / (v * v)                                      // Kinetic energy

	mpz := m / float64(z)

	fmt.Println(molecule.Mass(), mpz)
}

func (d Detector) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(int(d.Rect.Width()*float64(PXPM)), int(d.Rect.Height()*float64(PXPM)))
	img.Fill(color.RGBA{60, 75, 75, 255})

	drawOps := ebiten.DrawImageOptions{}
	drawOps.GeoM.Translate(float64(d.Rect.Min.X*float64(PXPM)), float64(d.Rect.Min.Y*float64(PXPM)))

	screen.DrawImage(img, &drawOps)
}

type Simulation struct {
	accelerationRegion ElectricField
	detector           Detector

	ionisationButton Button

	methane         Molecule
	drawableMethane RenderMolecule
}

func NewSimulation() *Simulation {
	s := &Simulation{
		accelerationRegion: ElectricField{
			Rect: NewRect(
				float64(100/PXPM), float64(150/PXPM),
				float64(300/PXPM), float64(750/PXPM),
			),
			PotentialDifference: 16_000,
		},
		detector: Detector{
			Rect: NewRect(
				float64(1400/PXPM), float64(150/PXPM),
				float64(1500/PXPM), float64(750/PXPM),
			),
		},

		ionisationButton: Button{
			Text:  "Ionise",
			Rect:  NewRect(100, 800, 300, 850),
			Color: color.RGBA{0, 70, 25, 255},
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
			Charge:     0,
			Pos:        Vec2{float64(100 / PXPM), float64(450 / PXPM)},
			Vel:        Vec2{0, 0},
			DriftTicks: 0,
		},
	}

	s.ionisationButton.Fuction = s.IoniseMolecules

	s.drawableMethane = RenderMolecule{
		&s.methane,
		color.RGBA{200, 210, 210, 255},
	}
	s.detector.AcellerationField = s.accelerationRegion

	return s
}

func (s *Simulation) IoniseMolecules() {
	for _, m := range []*Molecule{&s.methane} {
		if s.accelerationRegion.Rect.Contains(m.Pos) {
			m.Charge = 1
		}
	}
}

func (s *Simulation) Update() error {
	s.ionisationButton.Update()

	molecules := []*Molecule{}

	if s.methane.Active {
		s.methane.Update(s.accelerationRegion)
		molecules = append(molecules, &s.methane)
	}

	s.detector.Update(molecules)

	return nil
}

func (s Simulation) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 100, 120, 255})

	s.ionisationButton.Draw(screen)

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
