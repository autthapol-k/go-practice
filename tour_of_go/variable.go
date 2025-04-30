package main

import "fmt"

var i, j int = 1, 2

func Variable_WithInitialized() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func Variable_ShortDeclarations() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

var c, python, java bool

func TryVariable() {
	var i int
	fmt.Println(i, c, python, java)

	Variable_WithInitialized()
	Variable_ShortDeclarations()
}
