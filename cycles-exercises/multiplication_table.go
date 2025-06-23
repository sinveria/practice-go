package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number from 1 to 10: ")
	fmt.Scan(&num)

	if num < 1 || num > 10 {
		fmt.Println("Error: number must be between 1 and 10")
		return
	}

	fmt.Printf("\nMultiplication table for %d:\n", num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}

	fmt.Println("\nComplete multiplication table:")
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}
