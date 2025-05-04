package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Split(s, " ")
	for _, word := range words {
		// count := 0
		// for j := range len(words) {
		// 	if word == words[j] {
		// 		count++
		// 	}
		// }
		// m[word] = count
		m[word]++
	}

	return m
}

func TryExerciseMaps() {
	wc.Test(WordCount)
}
