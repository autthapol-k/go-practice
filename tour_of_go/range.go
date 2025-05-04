package main

import "fmt"

var pow_2 = []int{1, 2, 4, 8, 16, 32, 64, 128}

func TryRange() {
	for i, v := range pow_2 {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
