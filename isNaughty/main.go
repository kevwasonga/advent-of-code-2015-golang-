package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var niceCount int

	for scanner.Scan() {
		line := scanner.Text()
		if isNice(line) {
			niceCount++
		}
	}

	fmt.Println("Nice strings count:", niceCount)
}

func isNice(s string) bool {
	return hasRepeatedPair(s) && hasRepeatingWithOneBetween(s)
}

func hasRepeatedPair(s string) bool {
	pairMap := make(map[string]int)

	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		pairMap[pair]++

		// Check for overlapping pairs
		if pairMap[pair] == 2 && i > 0 && s[i-1] == s[i] {
			return true
		}
	}

	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if pairMap[pair] >= 2 {
			return true
		}
	}
	return false
}

func hasRepeatingWithOneBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
