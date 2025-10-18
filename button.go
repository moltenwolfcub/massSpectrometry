package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	mplusFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(ButtonFont))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

type Button struct {
	Text    string
	Rect    Rect
	Color   color.Color
	Fuction func()
}

func (b *Button) Update() {
	rawCursorX, rawCursorY := ebiten.CursorPosition()
	cursor := Vec2{float64(rawCursorX), float64(rawCursorY)}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && b.Rect.Contains(cursor) {
		b.Fuction()
		//TODO maybe change button color when pressed. maybe even hover color
	}
}

func (b Button) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(int(b.Rect.Width()), int(b.Rect.Height()))
	img.Fill(b.Color)

	op := &text.DrawOptions{}
	op.GeoM.Translate(0, 0) //TODO centre text
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(img, b.Text, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}, op)

	drawOps := ebiten.DrawImageOptions{}
	drawOps.GeoM.Translate(b.Rect.Min.X, b.Rect.Min.Y)

	screen.DrawImage(img, &drawOps)
}
