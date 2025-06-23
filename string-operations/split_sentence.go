package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Go is an open source programming language that makes it simple to build secure, scalable systems"

	words := strings.Fields(sentence)

	fmt.Println("Splitting a sentence into words:")
	for i, word := range words {
		fmt.Printf("%d: %s\n", i+1, word)
	}

	data := "apple,banana,orange"
	fruits := strings.Split(data, ",")
	fmt.Println("\nSplitting a string on commas:")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
