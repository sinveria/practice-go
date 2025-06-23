package main

import "fmt"

func main() {
	var firstNumber, secondNumber int

	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)

	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)

	fmt.Printf("%d + %d = %d\n", firstNumber, secondNumber, firstNumber+secondNumber)
	fmt.Printf("%d - %d = %d\n", firstNumber, secondNumber, firstNumber-secondNumber)
	fmt.Printf("%d * %d = %d\n", firstNumber, secondNumber, firstNumber*secondNumber)

	if secondNumber != 0 {
		fmt.Printf("%d / %d = %d\n", firstNumber, secondNumber, firstNumber/secondNumber)
		fmt.Printf("%d %% %d = %d\n", firstNumber, secondNumber, firstNumber%secondNumber)
	} else {
		fmt.Println("Division by zero is impossible.")
	}

	firstNumber++
	secondNumber--
	fmt.Printf("After increment: a = %d, b = %d\n", firstNumber, secondNumber)

	fmt.Printf("%d == %d: %t\n", firstNumber, secondNumber, firstNumber == secondNumber)
	fmt.Printf("%d != %d: %t\n", firstNumber, secondNumber, firstNumber != secondNumber)
	fmt.Printf("%d > %d: %t\n", firstNumber, secondNumber, firstNumber > secondNumber)
	fmt.Printf("%d < %d: %t\n", firstNumber, secondNumber, firstNumber < secondNumber)
}
