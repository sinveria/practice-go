package main

import "fmt"

func removeElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	colors := []string{"Red", "Green", "Blue", "Yellow", "Black"}

	fmt.Println("Original slice:", colors)

	colors = removeElement(colors, 2)
	fmt.Println("After removal:", colors)

	colors = removeElement(colors, 0)
	fmt.Println("After removing the first element:", colors)

	colors = removeElement(colors, len(colors)-1)
	fmt.Println("After removing the last element:", colors)
}
