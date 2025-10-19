package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var SELECTION []struct {
	m Molecule
	c color.Color
} = []struct {
	m Molecule
	c color.Color
}{
	{
		m: Molecule{
			Name: "methane",
			Atoms: []struct {
				element *Atom
				count   int
			}{{&CARBON, 1}, {&HYDROGEN, 4}},
		},
		c: color.RGBA{250, 50, 50, 255},
	},
}

type Selectable struct {
	DrawRegion Rect
	molecule   *Molecule
	renderable RenderMolecule
}

type Selector struct {
	Simulation *Simulation

	Rect    Rect
	Options []Selectable
}

func NewSelector(Rect Rect, simulation *Simulation) Selector {
	s := Selector{
		Simulation: simulation,
		Rect:       Rect,
	}

	tileSize := s.Rect.Height()

	PADDING := 5

	count := 0
	for _, o := range SELECTION {
		s.Options = append(s.Options, Selectable{
			DrawRegion: NewRect((tileSize+float64(PADDING))*float64(count)+float64(PADDING), float64(PADDING), (tileSize+float64(PADDING))*float64(count)+tileSize, tileSize-float64(PADDING)),
			molecule:   &o.m,
			renderable: RenderMolecule{
				Molecule: &o.m,
				Color:    o.c,
			},
		})
	}

	return s
}

func (s Selector) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		rawCursorX, rawCursorY := ebiten.CursorPosition()
		cursor := Vec2{float64(rawCursorX), float64(rawCursorY)}
		for _, o := range s.Options {
			if o.DrawRegion.Contains(cursor) {
				newMolecule := *o.molecule
				newMolecule.Active = true
				newMolecule.Charge = 0
				newMolecule.Vel = Vec2{0, 0}
				newMolecule.Pos = s.Simulation.GetSpawn()
				newMolecule.DriftTicks = 0

				s.Simulation.molecules = append(s.Simulation.molecules, &newMolecule)
				s.Simulation.drawableMolecules = append(s.Simulation.drawableMolecules, RenderMolecule{
					Molecule: &newMolecule,
					Color:    o.renderable.Color,
				})
				break
			}
		}
	}
}

func (s Selector) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(int(s.Rect.Width()), int(s.Rect.Height()))
	img.Fill(color.RGBA{170, 170, 170, 255})

	drawOps := ebiten.DrawImageOptions{}
	drawOps.GeoM.Translate(s.Rect.Min.Elem())

	screen.DrawImage(img, &drawOps)

	for _, option := range s.Options {
		img := ebiten.NewImage(int(option.DrawRegion.Width()), int(option.DrawRegion.Height()))
		img.Fill(color.RGBA{200, 200, 200, 255})

		MOLECULE_RADIUS := 10.0

		moleculePos := option.DrawRegion.Min.Add(option.DrawRegion.Size().Mul(0.5)).Add(Vec2{-MOLECULE_RADIUS / 2, -MOLECULE_RADIUS / 2})
		vector.FillCircle(img, float32(moleculePos.X), float32(moleculePos.Y), float32(MOLECULE_RADIUS), option.renderable.Color, true)

		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(option.DrawRegion.Min.Elem())

		screen.DrawImage(img, &op)
	}

	rawCursorX, rawCursorY := ebiten.CursorPosition()
	cursor := Vec2{float64(rawCursorX), float64(rawCursorY)}

	FONT_SIZE := 32
	OFFSET := Vec2{15, 3}
	TEXT_COLOR := color.White
	PADDING := 5
	BG_COLOR := color.RGBA{70, 70, 70, 200}

	for _, option := range s.Options {
		if option.DrawRegion.Contains(cursor) {
			drawText := cases.Title(language.English).String(option.molecule.Name)

			w, h := text.Measure(drawText, &text.GoTextFace{
				Source: fontSource,
				Size:   float64(FONT_SIZE),
			}, 0)

			img := ebiten.NewImage(int(w+float64(PADDING*2)), int(h+float64(PADDING*2)))
			img.Fill(BG_COLOR)

			textOp := &text.DrawOptions{}
			textOp.GeoM.Translate(float64(PADDING), float64(PADDING))
			textOp.ColorScale.ScaleWithColor(TEXT_COLOR)

			text.Draw(img, drawText, &text.GoTextFace{
				Source: fontSource,
				Size:   float64(FONT_SIZE),
			}, textOp)

			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(cursor.Add(OFFSET).Elem())

			screen.DrawImage(img, &op)

			break
		}
	}
}
