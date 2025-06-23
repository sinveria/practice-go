package main

import "fmt"

func LongestString(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]
	for _, s := range strings[1:] {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	result := LongestString(
		"kiwi",
		"banana",
		"grapefruit",
		"lemon",
		"apple",
	)

	fmt.Println("The longest string is:", result)

	countries := []string{"Brazil", "Japan", "United States", "Germany"}
	longestInSlice := LongestString(countries...)
	fmt.Println("Longest country name:", longestInSlice)
}
