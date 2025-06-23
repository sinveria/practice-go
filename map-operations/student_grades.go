package main

import "fmt"

func main() {
	grades := make(map[string]int)

	grades["Alexander"] = 5
	grades["Sophie"] = 4
	grades["Margarita"] = 5

	moreGrades := map[string]int{
		"Eve":       3,
		"Stanislav": 4,
	}

	for k, v := range moreGrades {
		grades[k] = v
	}

	if grade, exists := grades["Alexander"]; exists {
		fmt.Printf("Grade Alexander: %d\n", grade)
	} else {
		fmt.Println("Student Alexander not found")
	}

	fmt.Println("Before delete Sophie:", grades)
	delete(grades, "Sophie")
	fmt.Println("After delete Sophie:", grades)

	fmt.Println("\nAll grades:")
	for name, grade := range grades {
		fmt.Printf("%s: %d\n", name, grade)
	}
}
