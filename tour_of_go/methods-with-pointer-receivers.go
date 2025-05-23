package main

import (
	"fmt"
	"math"
)

type Vertex9 struct {
	X, Y float64
}

func (v *Vertex9) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func (v *Vertex9) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func TryMethodsWithPointerReceivers() {
	v := &Vertex9{3, 4}
	// v := Vertex9{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())

	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
