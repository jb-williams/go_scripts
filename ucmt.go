package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Default comment prefix
	commentPrefix := "#"

	// Check if a custom prefix is provided as an argument
	if len(os.Args) > 1 {
		commentPrefix = os.Args[1]
	}

	// Compile the regex pattern for removing the comment prefix
	pattern := fmt.Sprintf(`^(\s*)(%s ?)?`, regexp.QuoteMeta(commentPrefix))
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error compiling regex:", err)
		os.Exit(1)
	}

	// Read input line by line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		uncommentedLine := regex.ReplaceAllString(line, "$1")
		fmt.Println(uncommentedLine)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
}
