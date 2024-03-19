package main

import (
	"fmt"
)

var alienValues = map[byte]int{
	'A': 1,
	'B': 5,
	'Z': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'R': 1000,
}

func alienToInt(s string) (int, string) {
	result := 0
	prev := 0
	explanation := "Explanation: "
	for i := len(s) - 1; i >= 0; i-- {
		value := alienValues[s[i]]
		if value < prev {
			result -= value
			explanation += fmt.Sprintf("%c%c = %d, ", s[i], s[i+1], prev-value)
		} else {
			result += value
			prev = value
			explanation += fmt.Sprintf("%c = %d, ", s[i], value)
		}
	}
	return result, explanation
}

func main() {
	var input string
	fmt.Print("Input: s = ")
	fmt.Scanln(&input)
	output, explanation := alienToInt(input)
	fmt.Printf("Output: %d \n%s\n", output, explanation)
}
