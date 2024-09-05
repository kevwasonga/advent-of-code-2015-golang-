package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var totalPaper int

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line and process the data
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line by 'x'
		parts := strings.Split(line, "x")
		if len(parts) != 3 {
			fmt.Println("Unexpected line format:", line)
			continue
		}

		// Convert parts to integers
		l, err1 := strconv.Atoi(parts[0])
		w, err2 := strconv.Atoi(parts[1])
		h, err3 := strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Error converting to integer:", err1, err2, err3)
			continue
		}

		// Calculate the surface area
		surfaceArea := 2*l*w + 2*w*h + 2*h*l
		// Calculate the area of each side
		side1 := l * w
		side2 := w * h
		side3 := h * l
		// Find the smallest side
		smallestSide := min(side1, min(side2, side3))
		// Calculate the total paper needed for this box
		totalPaper += surfaceArea + smallestSide
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Total square feet of wrapping paper needed: %d\n", totalPaper)
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
