package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Default comment prefix
	commentPrefix := "#"

	// Check if a custom prefix is provided as an argument
	if len(os.Args) > 1 {
		commentPrefix = os.Args[1]
	}

	// Read input line by line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s %s\n", commentPrefix, line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
}
