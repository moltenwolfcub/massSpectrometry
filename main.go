package main

import (
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
	Rect Rect
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
}

func NewSimulation() *Simulation {
	s := &Simulation{
		accelerationRegion: ElectricField{
			Rect:                NewRect(100, 150, 300, 750),
			PotentialDifference: 16_000,
		},
		detector: Detector{
			Rect: NewRect(1400, 150, 1500, 750),
		},

		methane: Molecule{
			Name: "methane",
			Atoms: []struct {
				element *Atom
				count   int
			}{
				{&CARBON, 1},
				{&HYDROGEN, 4},
			},
			Charge: 1,
			Pos:    Vec2{120, 450},
			Vel:    Vec2{0, 0},
		},
	}

	s.drawableMethane = RenderMolecule{
		&s.methane,
		color.RGBA{200, 210, 210, 255},
	}

	return s
}

func (s *Simulation) Update() error {

	s.methane.Update(s.accelerationRegion)

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
