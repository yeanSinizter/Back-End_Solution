package main

import (
	"fmt"
	"strings"
)

func decodeEncodedNumber(encoded string) string {
	// Initialize variables
	var result strings.Builder
	sum := 0
	leftValue := 0

	// Iterate through each character in the encoded string
	for _, char := range encoded {
		switch char {
		case 'L':
			// Check if the left value is greater than the right value
			if leftValue > 0 {
				result.WriteString(fmt.Sprintf("%d", leftValue))
			} else {
				result.WriteString(fmt.Sprintf("-%d", leftValue))
			}
			leftValue++
			sum += leftValue
		case 'R':
			// Check if the right value is greater than the left value
			if leftValue > 0 {
				result.WriteString(fmt.Sprintf("-%d", leftValue))
			} else {
				result.WriteString(fmt.Sprintf("%d", leftValue))
			}
			result.WriteString("0")
		case '=':
			result.WriteString("=")
		}
	}

	return result.String()
}

func main() {
	// Get input from the keyboard
	var encoded string
	fmt.Print("Enter the encoded text: ")
	fmt.Scanln(&encoded)

	// Decode the encoded text
	decoded := decodeEncodedNumber(encoded)

	// Print the decoded numerical sequence
	fmt.Println("Decoded numerical sequence:", decoded)
}
