package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = math.Pi
	const e = math.E

	circleArea := pi * math.Pow(5, 2)
	naturalLog := math.Log(e)

	fmt.Printf("π = %.4f\n", pi)
	fmt.Printf("e = %.4f\n", e)
	fmt.Printf("Area of a circle (r=5) = %.2f\n", circleArea)
	fmt.Printf("ln(e) = %.2f\n", naturalLog)
}
