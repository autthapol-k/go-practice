package main

import (
	"fmt"
	"math"
)

type Vertex5 struct {
	X, Y float64
}

func (v Vertex5) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex5) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func TryMethodsPointers() {
	v := Vertex5{3, 4}
	v.Scale(10)

	fmt.Println(v.Abs())
	fmt.Println(v)
}
