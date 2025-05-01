package main

import (
	"fmt"
	"math"
)

func pow_IfAndElse(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// cann't use v here, though
	return lim
}

func TryIfAndElse() {
	fmt.Println(
		pow_IfAndElse(3, 2, 10),
		pow_IfAndElse(3, 3, 20),
	)
}
