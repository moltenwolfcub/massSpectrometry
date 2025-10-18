package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Simulation struct {
	accelerationRegion Rect

	methane         Molecule
	drawableMethane RenderMolecule
}

func NewSimulation() *Simulation {
	s := &Simulation{
		accelerationRegion: NewRect(200, 150, 500, 750),

		methane: Molecule{
			Name: "methane",
			Atoms: []struct {
				element *Atom
				count   int
			}{
				{&CARBON, 1},
				{&HYDROGEN, 4},
			},
			Charge: 0,
			Pos:    Vec2{250, 450},
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

	accelRegion := ebiten.NewImage(int(s.accelerationRegion.Width()), int(s.accelerationRegion.Height()))
	accelRegion.Fill(color.RGBA{250, 50, 50, 100})
	drawOps := ebiten.DrawImageOptions{}
	drawOps.GeoM.Translate(s.accelerationRegion.Min.X, s.accelerationRegion.Min.Y)
	screen.DrawImage(accelRegion, &drawOps)

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
