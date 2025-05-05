package main

import (
	"fmt"
	"math"
)

type Vertex8 struct {
	X, Y float64
}

func (v Vertex8) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex8) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func TryIndirectionValues() {
	v := Vertex8{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex8{4, 3}
	fmt.Println(p.Abs()) // (*p).Abs()
	fmt.Println(AbsFunc(*p))
}
