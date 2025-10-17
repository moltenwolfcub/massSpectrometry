package main

import (
	"fmt"
)

type Vec2 struct {
	X, Y float32
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%v,%v)", v.X, v.Y)
}

func (v Vec2) Elem() (float32, float32) {
	return v.X, v.Y
}
