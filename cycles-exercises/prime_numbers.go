package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Prime numbers up to 100:")
	count := 0
	for num := 2; num <= 100; num++ {
		if isPrime(num) {
			fmt.Printf("%4d", num)
			count++
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}
}
