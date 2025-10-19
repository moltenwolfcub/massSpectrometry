package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ButtonColor struct {
	Primary   color.Color
	Secondary color.Color
	Hover     color.Color
}

func (c ButtonColor) GetPrimary() color.Color {
	if c.Primary == nil {
		panic("No primary color set.")
	}
	return c.Primary
}

func (c ButtonColor) GetHover() color.Color {
	if c.Hover == nil {
		return c.GetPrimary()
	}
	return c.Hover
}

func (c ButtonColor) GetSecondary() color.Color {
	if c.Secondary == nil {
		return c.GetHover()
	}
	return c.Secondary
}

type buttonState int

const (
	normal buttonState = iota
	clicked
	hover
)

func (b buttonState) getColor(c ButtonColor) color.Color {
	switch b {
	case normal:
		return c.GetPrimary()
	case clicked:
		return c.GetSecondary()
	case hover:
		return c.GetHover()
	default:
		fmt.Println("[Warning] Unknown button state. Returning primary color.")
		return c.GetPrimary()
	}
}

type Button struct {
	Text         string
	TextColor    ButtonColor
	TextSize     float64
	Rect         Rect
	ButtonColor  ButtonColor
	Fuction      func()
	state        buttonState
	clickTime    int
	MaxClickTime int
}

func (b *Button) Update() {
	if b.state == clicked {
		b.clickTime++
	} else {
		b.clickTime = 0
	}

	rawCursorX, rawCursorY := ebiten.CursorPosition()
	cursor := Vec2{float64(rawCursorX), float64(rawCursorY)}
	if b.Rect.Contains(cursor) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			b.state = clicked
			b.Fuction()
		} else {
			if b.state == clicked {
				if b.clickTime > b.MaxClickTime {
					b.state = hover
				}
			} else {
				b.state = hover
			}
		}
	} else {
		if b.state == clicked {
			if b.clickTime > b.MaxClickTime {
				b.state = normal
			}
		} else {
			b.state = normal
		}
	}
}

func (b Button) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(int(b.Rect.Width()), int(b.Rect.Height()))
	img.Fill(b.state.getColor(b.ButtonColor))

	w, h := text.Measure(b.Text, &text.GoTextFace{
		Source: fontSource,
		Size:   b.TextSize,
	}, 0)

	op := &text.DrawOptions{}
	op.GeoM.Translate((b.Rect.Width()-w)/2, (b.Rect.Height()-h)/2)
	op.ColorScale.ScaleWithColor(b.state.getColor(b.TextColor))

	text.Draw(img, b.Text, &text.GoTextFace{
		Source: fontSource,
		Size:   b.TextSize,
	}, op)

	drawOps := ebiten.DrawImageOptions{}
	drawOps.GeoM.Translate(b.Rect.Min.X, b.Rect.Min.Y)

	screen.DrawImage(img, &drawOps)
}
