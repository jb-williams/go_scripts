package main

import (
	"fmt"
	"os"
	"strings"
)

// Function to create beginning dashes
func createDashes(text string) string {
	textLength := len(text)
	dashCount := (76 - textLength) / 2
	return strings.Repeat("-", dashCount)
}

// Function to create ending dashes
func createEndDashes(dashes, text string) string {
	dashLength := len(dashes)
	textLength := len(text)
	endDashCount := 76 - (textLength + dashLength)
	return strings.Repeat("-", endDashCount)
}

// Function to create the header format
func createHeader(text string) string {
	dashes := createDashes(text)
	endDashes := createEndDashes(dashes, text)
	return fmt.Sprintf("# %s %s %s #", dashes, text, endDashes)
}

// Main function
func main() {
	var inputText string

	// Check if arguments are provided or read from STDIN
	if len(os.Args) > 1 {
		inputText = strings.Join(os.Args[1:], " ")
	} else {
		// Read from STDIN if no arguments are given
		fmt.Scanln(&inputText)
	}

	// Generate and print the header
	header := createHeader(inputText)
	fmt.Println(header)
}

