package main

import "fmt"

func Sqrt(x float64) float64 {
	count := 1
	// z := float64(1)
	z := x / 2
	for count <= 10 {
		old := z
		z -= (z*z - x) / (2 * z)
		if old == z {
			// fmt.Println(z)
			return z
		}

		count++
		fmt.Println(z)
	}
	return z
}

func TryExerciseLoopAndFunctions() {
	// fmt.Println(Sqrt(2))
	Sqrt(5)
}
