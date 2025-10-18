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
	}
}

func (r Rect) Width() float64 {
	return math.Abs(r.Max.X - r.Min.X)
}
func (r Rect) Height() float64 {
	return math.Abs(r.Max.Y - r.Min.Y)
}
