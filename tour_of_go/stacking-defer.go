package main

import "fmt"

func TryStackingDefer() {
	fmt.Println("Counting...")

	for i := range 10 {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
