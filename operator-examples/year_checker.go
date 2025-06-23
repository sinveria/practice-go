package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func main() {
	var year int
	fmt.Print("Enter year: ")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Printf("%d year - leap year\n", year)
	} else {
		fmt.Printf("%d the year is not a leap year\n", year)
	}
}
