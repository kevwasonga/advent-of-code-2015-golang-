package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the file containing the input strings
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize the count of nice strings
	niceCount := 0

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the current line
		line := strings.TrimSpace(scanner.Text())

		// Check if the current string is "nice"
		if isNice(line) {
			niceCount++
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Output the count of nice strings
	fmt.Printf("Total number of nice strings: %d\n", niceCount)
}

// Function to check if a string is "nice"
func isNice(s string) bool {
	vowelCount := 0
	double := false
	forbidden := false

	// Check for vowels and double letters
	for i := 0; i < len(s); i++ {
		// Check for vowels
		if isVowel(s[i]) {
			vowelCount++
		}

		// Check for double letters (only if i + 1 is a valid index)
		if i < len(s)-1 && s[i] == s[i+1] {
			double = true
		}

		// Check for forbidden substrings (only if i + 1 is a valid index)
		if i < len(s)-1 && (s[i] == 'a' && s[i+1] == 'b' || s[i] == 'c' && s[i+1] == 'd' || s[i] == 'p' && s[i+1] == 'q' || s[i] == 'x' && s[i+1] == 'y') {
			forbidden = true
		}
	}

	// The string is "nice" if it has at least 3 vowels, a double letter, and no forbidden substrings
	return vowelCount >= 3 && double && !forbidden
}

// Helper function to check if a character is a vowel
func isVowel(r byte) bool {
	vowels := "aeiouAEIOU"
	for i := 0; i < len(vowels); i++ {
		if r == vowels[i] {
			return true
		}
	}
	return false
}
