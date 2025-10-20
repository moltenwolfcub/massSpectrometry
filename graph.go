package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Graph struct {
	Sim     *Simulation
	Data    *DataLogger
	Rect    Rect
	Buttons []*Button
}

func NewGraph(sim *Simulation) Graph {
	g := Graph{
		Sim:  sim,
		Data: &sim.detector.DataLogger,
		Rect: NewRect(50, 50, 1550, 750),
	}

	g.Buttons = []*Button{
		{
			Text: "Back",
			TextColor: ButtonColor{
				Primary: color.White,
			},
			TextSize: 30,
			Rect:     NewRect(1350, 800, 1500, 850),
			ButtonColor: ButtonColor{
				Primary:   color.RGBA{0, 70, 25, 255},
				Hover:     color.RGBA{0, 63, 22, 255},
				Secondary: color.RGBA{0, 94, 35, 255},
			},
			Fuction:      g.Exit,
			MaxClickTime: 10,
		},
	}

	return g
}

func (g *Graph) Exit() {
	g.Sim.currentScreen = MainScreen
}

func (g *Graph) Update() {
	for _, b := range g.Buttons {
		b.Update()
	}
}

func (g *Graph) Draw(screen *ebiten.Image) {
	//graph
	img := ebiten.NewImage(int(g.Rect.Width()), int(g.Rect.Height()))
	img.Fill(color.RGBA{170, 170, 170, 255})
	PAD := 10.0
	TEXT_PAD := 5.0
	LABEL_X := "m/z"
	LABEL_Y := "Relative Abundance"
	FONT_SIZE := 36.0
	LABEL_COLOR := color.Black
	AXIS_THICKNESS := float32(20)
	AXIS_COLOR := color.Black

	wx, hx := text.Measure(LABEL_X, &text.GoTextFace{
		Source: fontSource,
		Size:   FONT_SIZE,
	}, 0)
	wy, hy := text.Measure(LABEL_Y, &text.GoTextFace{
		Source: fontSource,
		Size:   FONT_SIZE,
	}, 0)

	axisRect := NewRect(
		PAD+hy+TEXT_PAD,
		PAD,
		g.Rect.Width()-PAD,
		g.Rect.Height()-PAD-hx-TEXT_PAD,
	)

	//graph:labels
	xLabelOps := &text.DrawOptions{}
	xLabelOps.GeoM.Translate(axisRect.Min.X+(axisRect.Width()-wx)/2, g.Rect.Height()-PAD-hx)
	xLabelOps.ColorScale.ScaleWithColor(LABEL_COLOR)

	text.Draw(img, LABEL_X, &text.GoTextFace{
		Source: fontSource,
		Size:   FONT_SIZE,
	}, xLabelOps)

	yLabelOps := &text.DrawOptions{}
	yLabelOps.GeoM.Rotate(-math.Pi / 2)
	yLabelOps.GeoM.Translate(PAD, axisRect.Min.Y+(axisRect.Height()-wy)/2+wy)
	yLabelOps.ColorScale.ScaleWithColor(LABEL_COLOR)

	text.Draw(img, LABEL_Y, &text.GoTextFace{
		Source: fontSource,
		Size:   FONT_SIZE,
	}, yLabelOps)

	//graph:axis
	axisImg := ebiten.NewImage(int(axisRect.Width()), int(axisRect.Height()))

	vector.StrokeLine(axisImg, 0, float32(axisRect.Height())-AXIS_THICKNESS/2, float32(axisRect.Width()), float32(axisRect.Height())-AXIS_THICKNESS/2, AXIS_THICKNESS, AXIS_COLOR, true)
	vector.StrokeLine(axisImg, AXIS_THICKNESS/2, 0, AXIS_THICKNESS/2, float32(axisRect.Height()), AXIS_THICKNESS, AXIS_COLOR, true)

	axisOps := ebiten.DrawImageOptions{}
	axisOps.GeoM.Translate(axisRect.Min.Elem())
	img.DrawImage(axisImg, &axisOps)

	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Translate(g.Rect.Min.Elem())
	screen.DrawImage(img, &ops)

	// buttons
	for _, b := range g.Buttons {
		b.Draw(screen)
	}
}
