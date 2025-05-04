package main

import "fmt"

// Go has pointers. A pointer holds the memory address of a value.
func TryPointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through to the pointer
	*p = 21         // set i through to the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 7    // divided j through the pointer
	fmt.Println(j) // see the value of j
}
