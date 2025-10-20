package main

import (
	"cmp"
	"fmt"
	"image/color"
	"math"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	for _, b := range g.Buttons {
		b.state = normal
	}
}

func (g *Graph) Update() {
	for _, b := range g.Buttons {
		b.Update()
	}
}

func (g Graph) Draw(screen *ebiten.Image) {
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
	AXIS_COLOR := color.RGBA{50, 50, 50, 255}

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

	drawTooltip := g.drawData(axisImg, axisRect, float64(AXIS_THICKNESS), img, axisRect.Min.Add(g.Rect.Min))

	axisOps := ebiten.DrawImageOptions{}
	axisOps.GeoM.Translate(axisRect.Min.Elem())
	img.DrawImage(axisImg, &axisOps)

	if drawTooltip != nil {
		drawTooltip(img, g.Rect.Min)
	}

	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Translate(g.Rect.Min.Elem())
	screen.DrawImage(img, &ops)

	// buttons
	for _, b := range g.Buttons {
		b.Draw(screen)
	}
}

func (g Graph) drawData(img *ebiten.Image, rect Rect, axisThickness float64, tooltipImg *ebiten.Image, topLeft Vec2) (drawTooltip func(*ebiten.Image, Vec2)) {
	if len(g.Data.data) == 0 {
		return
	}

	END_FILLER := 10.0
	TOP_PAD := 20.0
	LINE_WIDTH := 10.0
	LINE_COLOR := color.RGBA{50, 50, 50, 255}
	TOOLTIP_FONT_SIZE := 32.0
	TOOLTIP_OFFSET := Vec2{17, 3}
	TOOLTIP_TEXT_COLOR := color.White
	TOOLTIP_PADDING := 5
	TOOLTIP_BG_COLOR := color.RGBA{70, 70, 70, 200}
	TOOLTIP_LINE_SPACING := 5.0

	largestMZ := slices.MaxFunc(g.Data.data, func(a, b LoggerEntry) int {
		return cmp.Compare(a.mz, b.mz)
	})

	largestAbundance := slices.MaxFunc(g.Data.data, func(a, b LoggerEntry) int {
		return cmp.Compare(a.abundance, b.abundance)
	})

	maxX := largestMZ.mz + END_FILLER

	rawCursorX, rawCursorY := ebiten.CursorPosition()
	cursor := Vec2{float64(rawCursorX), float64(rawCursorY)}

	for _, e := range g.Data.data {
		percentAlong := e.mz / maxX
		horizontalPos := percentAlong * (rect.Width() - axisThickness)

		percentHeight := float64(e.abundance) / float64(largestAbundance.abundance)
		height := percentHeight * (rect.Height() - axisThickness - TOP_PAD)

		vector.StrokeLine(img,
			float32(horizontalPos), float32(rect.Height()-axisThickness),
			float32(horizontalPos), float32(rect.Height()-axisThickness-height),
			float32(LINE_WIDTH), LINE_COLOR, true,
		)

		barRect := NewRect(
			horizontalPos-axisThickness/2.0,
			rect.Height()-axisThickness-height,
			horizontalPos+axisThickness/2.0,
			rect.Height()-axisThickness,
		).Translate(topLeft)
		if barRect.Contains(cursor) {
			tooltip := fmt.Sprintf("m/z: %-6.6f\nAbundance:%-6.3f%%", e.mz, float64(e.abundance)/float64(g.Data.totalEntries)*100)

			drawText := cases.Title(language.English).String(tooltip)

			w, h := text.Measure(drawText, &text.GoTextFace{
				Source: fontSource,
				Size:   TOOLTIP_FONT_SIZE,
			}, TOOLTIP_FONT_SIZE+TOOLTIP_LINE_SPACING)

			textImg := ebiten.NewImage(int(w+float64(TOOLTIP_PADDING*2)), int(h+float64(TOOLTIP_PADDING*2)))
			textImg.Fill(TOOLTIP_BG_COLOR)

			textOp := &text.DrawOptions{}
			textOp.GeoM.Translate(float64(TOOLTIP_PADDING), float64(TOOLTIP_PADDING))
			textOp.ColorScale.ScaleWithColor(TOOLTIP_TEXT_COLOR)
			textOp.LayoutOptions.LineSpacing = TOOLTIP_FONT_SIZE + TOOLTIP_LINE_SPACING

			text.Draw(textImg, drawText, &text.GoTextFace{
				Source: fontSource,
				Size:   float64(TOOLTIP_FONT_SIZE),
			}, textOp)

			// tooltipImg.DrawImage(textImg, &op)
			drawTooltip = func(i *ebiten.Image, iTopLeft Vec2) {
				op := ebiten.DrawImageOptions{}
				op.GeoM.Translate(cursor.Sub(iTopLeft).Add(TOOLTIP_OFFSET).Elem())
				i.DrawImage(textImg, &op)
			}
		}
	}
	return
}
