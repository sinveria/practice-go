package main

import "fmt"

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 10
	y := 20

	fmt.Printf("Before exchange: x = %d, y = %d\n", x, y)

	Swap(&x, &y)

	fmt.Printf("After the exchange: x = %d, y = %d\n", x, y)

	str1 := "hello"
	str2 := "world"

	fmt.Printf("\nBefore exchanging strings: s1 = %s, s2 = %s\n", str1, str2)
	swapStrings(&str1, &str2)
	fmt.Printf("After exchanging strings: s1 = %s, s2 = %s\n", str1, str2)
}

func swapStrings(a, b *string) {
	*a, *b = *b, *a
}
