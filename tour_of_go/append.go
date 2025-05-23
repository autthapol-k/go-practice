package main

func TryAppend() {
	var s []int
	printSlice(s)

	// append works on nil slices
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more then one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}
