package main

import (
	"fmt"
)

func main() {
	var dayNum int
	fmt.Print("Enter the day of the week number (1-7): ")
	fmt.Scan(&dayNum)

	var dayName string

	switch dayNum {
	case 1:
		dayName = "Monday"
	case 2:
		dayName = "Tuesday"
	case 3:
		dayName = "Wednesday"
	case 4:
		dayName = "Thursday"
	case 5:
		dayName = "Friday"
	case 6:
		dayName = "Saturday"
	case 7:
		dayName = "Sunday"
	default:
		fmt.Println("Error: enter a number between 1 and 7")
		return
	}

	fmt.Printf("Day of the week №%d - %s\n", dayNum, dayName)
}
