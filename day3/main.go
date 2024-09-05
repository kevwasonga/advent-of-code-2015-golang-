package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "^v^v^v^v^v" // Input string that represents movement along the x and y axes

	var coordinates [][]int // To store the movement vectors corresponding to each input character

	// Convert input characters to movement vectors
	for _, char := range input {
		switch char {
		case '^': // Move up (increase y)
			coordinates = append(coordinates, []int{0, 1})
		case 'v': // Move down (decrease y)
			coordinates = append(coordinates, []int{0, -1})
		case '>': // Move right (increase x)
			coordinates = append(coordinates, []int{1, 0})
		case '<': // Move left (decrease x)
			coordinates = append(coordinates, []int{-1, 0})
		default:
			continue // Ignore other characters
		}
	}

	var homes [][]int               // To store the homes (coordinates) visited during the journey
	currentH := []int{0, 0}         // Starting point at coordinate (0, 0)
	homes = append(homes, currentH) // Add the starting home to the list of visited homes

	// Traverse the coordinates and calculate the homes visited
	for _, move := range coordinates {
		// Update the current home by adding the movement vector to the current position
		currentX := currentH[0] + move[0]
		currentY := currentH[1] + move[1]

		// Update the current home position
		currentH = []int{currentX, currentY}

		// Add the new home (coordinate) to the homes visited
		homes = append(homes, currentH)
	}

	// Perform the "cleanhomes operation" to remove duplicate homes
	cleanHomes := removeDuplicates(homes)

	// Print the original homes (with duplicates)
	fmt.Println("Original Homes:")
	fmt.Println(homes)

	// Print the clean homes (without duplicates)
	fmt.Println("Clean Homes:")
	fmt.Println(cleanHomes)

	// Print the number of unique homes
	fmt.Println("Number of unique homes:")
	fmt.Println(len(cleanHomes))
}

// Helper function to remove duplicate coordinates from the homes slice
func removeDuplicates(homes [][]int) [][]int {
	// Create a map to store unique coordinates as strings ("x,y")
	uniqueHomes := make(map[string]bool)
	var result [][]int

	// Iterate over each home in the homes slice
	for _, home := range homes {
		// Convert the coordinate to a string in the format "x,y"
		key := strconv.Itoa(home[0]) + "," + strconv.Itoa(home[1])

		// Check if this coordinate (key) has already been seen
		if !uniqueHomes[key] {
			// If the coordinate hasn't been seen, add it to the result slice
			result = append(result, home)
			// Mark the coordinate as seen by adding it to the map
			uniqueHomes[key] = true
		}
	}

	// Return the result, which contains only unique homes
	return result
}
