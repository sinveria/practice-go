package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go Programming Language"

	fmt.Printf("Source string: \"%s\"\n", text)
	fmt.Printf("Line length: %d characters\n", len(text))

	substr := "Prog"
	fmt.Printf("Contains a substring \"%s\": %t\n", substr, strings.Contains(text, substr))

	fmt.Printf("Upper case: %s\n", strings.ToUpper(text))
	fmt.Printf("Lower case: %s\n", strings.ToLower(text))

	fmt.Printf("Replacement: %s\n", strings.ReplaceAll(text, "Go", "Golang"))
	fmt.Printf("Number of letters 'g': %d\n", strings.Count(text, "g"))
}
