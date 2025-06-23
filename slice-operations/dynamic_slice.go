package main

import "fmt"

func main() {
	var fruits []string

	fruits = append(fruits, "Qiwi")
	fruits = append(fruits, "Lemon", "Orange")

	moreFruits := []string{"Apple", "Mango"}
	fruits = append(fruits, moreFruits...)

	fmt.Println("Slice fruits:", fruits)
	fmt.Println("Lenght:", len(fruits))
	fmt.Println("Capacity:", cap(fruits))

	fmt.Println("\nSlice elements:")
	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i, fruit)
	}
}
