package main

import (
	"fmt"
	"math"
)

type Vertex6 struct {
	X, Y float64
}

func Abs2(v Vertex6) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v Vertex6, f float64) {
	v.X *= f
	v.Y *= f
}

func TryMethodsPointersExplained() {
	v := Vertex6{3, 4}
	Scale(v, 10)
	fmt.Println(Abs2(v))
}
