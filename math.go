package main

import (
	"fmt"
	"math"
)

type Vec2 struct {
	X, Y float64
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%v,%v)", v.X, v.Y)
}

func (v Vec2) Elem() (float64, float64) {
	return v.X, v.Y
}

func (v Vec2) Add(u Vec2) Vec2 {
	return Vec2{v.X + u.X, v.Y + u.Y}
}

func (v Vec2) Sub(u Vec2) Vec2 {
	return Vec2{v.X - u.X, v.Y - u.Y}
}

func (v Vec2) Mul(u float64) Vec2 {
	return Vec2{v.X * u, v.Y * u}
}

type Rect struct {
	Min Vec2
	Max Vec2
}

func NewRect(minX, minY, maxX, maxY float64) Rect {
	return Rect{
		Min: Vec2{
			X: minX,
			Y: minY,
		},
		Max: Vec2{
			X: maxX,
			Y: maxY,
		},
	}.Canon()
}

func (r Rect) Width() float64 {
	return math.Abs(r.Max.X - r.Min.X)
}
func (r Rect) Height() float64 {
	return math.Abs(r.Max.Y - r.Min.Y)
}
func (r Rect) Size() Vec2 {
	return Vec2{r.Width(), r.Height()}
}

// Copied straight from image.Point.In()
func (r Rect) Contains(v Vec2) bool {
	return r.Min.X <= v.X && v.X < r.Max.X && r.Min.Y <= v.Y && v.Y < r.Max.Y
}

// Makes sure the minimum and maximum points are as such.
// Copied straight from image.Rectangle.Canon()
func (r Rect) Canon() Rect {
	if r.Max.X < r.Min.X {
		r.Min.X, r.Max.X = r.Max.X, r.Min.X
	}
	if r.Max.Y < r.Min.Y {
		r.Min.Y, r.Max.Y = r.Max.Y, r.Min.Y
	}
	return r
}
