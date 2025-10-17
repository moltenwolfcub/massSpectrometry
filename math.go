package main

import (
	"fmt"
)

func Abs32(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}

type Vec2 struct {
	X, Y float32
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%v,%v)", v.X, v.Y)
}

func (v Vec2) Elem() (float32, float32) {
	return v.X, v.Y
}

type Rect struct {
	Min Vec2
	Max Vec2
}

func NewRect(minX, minY, maxX, maxY float32) Rect {
	return Rect{
		Min: Vec2{
			X: minX,
			Y: minY,
		},
		Max: Vec2{
			X: maxX,
			Y: maxY,
		},
	}
}

func (r Rect) Width() float32 {
	return Abs32(r.Max.X - r.Min.X)
}
func (r Rect) Height() float32 {
	return Abs32(r.Max.Y - r.Min.Y)
}
