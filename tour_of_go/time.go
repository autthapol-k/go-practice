package main

import (
	"fmt"
	"time"
)

func TryTimeFormat() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now().Format(time.RFC3339))
}
