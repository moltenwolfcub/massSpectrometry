package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	methane         Molecule
	drawableMethane RenderMolecule
}

func NewGame() *Game {
	g := &Game{
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
			Pos:    Vec2{800, 450},
			Vel:    Vec2{0, 0},
		},
	}

	g.drawableMethane = RenderMolecule{
		&g.methane,
		color.RGBA{200, 210, 210, 255},
	}

	return g
}

func (g Game) Update() error {
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 100, 120, 255})

	g.drawableMethane.Draw(screen)
}

func (g Game) Layout(actualWidth, actualHeight int) (screenWidth, screenHeight int) {
	return 1600, 900
}

func main() {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Mass Spectrometry")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
