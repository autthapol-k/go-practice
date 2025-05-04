package main

import "fmt"

func TryMakingSlices() {
	a := make([]int, 5)
	printSlice_2("a", a)

	b := make([]int, 0, 5)
	printSlice_2("b", b)

	c := b[:2]
	printSlice_2("c", c)

	d := c[1:5]
	printSlice_2("d", d)

	d[0] = 1
	fmt.Println(c)
	fmt.Println(d)
}

func printSlice_2(s string, x []int) {
	fmt.Printf("%s len=%d, cap=%d %v\n", s, len(x), cap(x), x)
}
