package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for _, greeting := range messages {
		fmt.Println(greeting)
	}
}
