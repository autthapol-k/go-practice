package main

import "fmt"

var m2 = map[string]Vertex2{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func TryMapLiterals() {
	fmt.Println(m2)
}
