package main

import (
	"fmt"
	"sort"
)

func Contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func SortSlice(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Ints(sorted)
	return sorted
}

func Filter(slice []int, condition func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	numbers := []int{5, 2, 8, 3, 1, 9}

	fmt.Println("Contains 8:", Contains(numbers, 8))
	fmt.Println("Contains 10:", Contains(numbers, 10))

	sorted := SortSlice(numbers)
	fmt.Println("Sorted:", sorted)

	even := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", even)
}
