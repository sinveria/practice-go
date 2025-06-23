package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	frequency := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		cleanWord := strings.Trim(lowerWord, ",.!?;:\"()")
		frequency[cleanWord]++
	}

	return frequency
}

func main() {
	text := "Go is an open source programming language that makes it simple to build secure, scalable systems."

	freq := wordFrequency(text)

	fmt.Println("Word frequency in text:")
	for word, count := range freq {
		fmt.Printf("%s: %d\n", word, count)
	}
}
