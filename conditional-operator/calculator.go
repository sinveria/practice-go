

import (
	"fmt"
	"os"
)

func main() {
	var firstNumber, secondNumber float64
	var op string

	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&op)

	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)

	var result float64
	var err error

	switch op {
	case "+":
		result = firstNumber + secondNumber
	case "-":
		result = firstNumber - secondNumber
	case "*":
		result = firstNumber * secondNumber
	case "/":
		if secondNumber == 0 {
			err = fmt.Errorf("division by zero is impossible")
		} else {
			result = firstNumber / secondNumber
		}
	default:
		err = fmt.Errorf("unknown operator: %s", op)
	}

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", firstNumber, op, secondNumber, result)
}
