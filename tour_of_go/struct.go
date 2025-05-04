package main

import "fmt"

// Struct is a collection of fields.
type Vertex struct {
	X int
	Y int
}

func TryStruct() {
	fmt.Print(Vertex{1, 2})
}
