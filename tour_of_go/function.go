package main

import "fmt"

func Add(x int, y int) int {
	return x + y
}

func Add_OmitType(x, y int) int {
	return x + y
}

func Swap_MultipleResult(x, y string) (string, string) {
	return y, x
}

func Split_NakedReturn(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func TryFunction() {
	fmt.Println(Add(2, 3))
	fmt.Println(Add_OmitType(2, 3))
	fmt.Println(Swap_MultipleResult("hello", "world"))
	fmt.Println(Split_NakedReturn(5))
}
