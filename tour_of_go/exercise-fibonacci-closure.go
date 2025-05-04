package main

import "fmt"

func fibonacci() func() int {
	x := 0
	var prev, curr int
	return func() int {
		if x == 0 {
			prev = 0
			curr = 0
			x++
			return 0
		}
		if x == 1 {
			prev = 0
			curr = 1
			x++
			return 1
		}

		temp := curr
		curr = prev + curr
		prev = temp

		x++
		return curr
	}
}

func TryExerciseFibonanciClosure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
