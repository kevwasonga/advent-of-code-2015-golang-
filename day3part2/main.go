package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "^>v<"
	// map sannas home movement directions
	var move1 string
	for i := 0; i < len(input); i += 2 {
		move1 += string(input[i])
	}
	// map robo sannas home movement directions
	var move2 string
	for i := 1; i < len(input); i += 2 {
		move2 += string(input[i])
	}
	// store sanna home coordinates

	var sannaCoordinates [][]int
	for _, char := range move1 {
		switch char {
		case '^':
			sannaCoordinates = append(sannaCoordinates, []int{0, 1})
		case 'v':
			sannaCoordinates = append(sannaCoordinates, []int{0, -1})
		case '>':
			sannaCoordinates = append(sannaCoordinates, []int{1, 0})
		case '<':
			sannaCoordinates = append(sannaCoordinates, []int{-1, 0})
		default:
			continue

		}
	}
	// coordinates where sanna will deliver gifts
	fmt.Println("These are the coordinates that sanna followed")
	fmt.Println(sannaCoordinates)

	var duplisanna [][]int
	for _, char := range move2 {
		switch char {
		case '^':
			duplisanna = append(duplisanna, []int{0, 1})
		case 'v':
			duplisanna = append(duplisanna, []int{0, -1})
		case '>':
			duplisanna = append(duplisanna, []int{1, 0})
		case '<':
			duplisanna = append(duplisanna, []int{-1, 0})
		default:
			continue

		}
	}
	// coordinates where dupicate sanna will deliver gifts
	fmt.Println("These are the coordinates that sannas robot followed")
	fmt.Println(duplisanna)

	// sum up sanna coordinates to identify the individual homes visited

	var homes1 [][]int
	current1 := []int{0, 0}
	homes1 = append(homes1, current1)

	for _, home := range sannaCoordinates {
		currentX := home[0] + current1[0]
		currentY := home[1] + current1[1]

		current1 = []int{currentX, currentY}
		homes1 = append(homes1, current1)
	}

	fmt.Println("These are the homes sanna visited and dropped gifts at:")
	fmt.Println(homes1)

	// find specific home coordinates for sannas duplicate gift delivery

	var homes2 [][]int

	current2 := []int{0, 0}

	for _, home := range duplisanna {

		currentX := current2[0] + home[0]
		currentY := current2[1] + home[1]

		current2 = []int{currentX, currentY}

		homes2 = append(homes2, current2)
	}
	fmt.Println("These are the homes ROBOsanna visited and dropped gifts at:")
	fmt.Println(homes2)

	// attempt to have a combined 2d slice of homes visited by both sanna and sannas robotgo

	for _, home := range homes2 {
		homes1 = append(homes1, home)
	}

	fmt.Println("combined homes visited by sanna and robo sanna")
	fmt.Println(homes1)

	// clean combined homes to remove duplicates
	combined := cleanHomes(homes1)
	// output number of homes visited with no duplicates
	fmt.Println(len(combined))
}

func cleanHomes(homes [][]int) [][]int {
	unique := make(map[string]bool)
	var result [][]int

	for _, home := range homes {
		// continously aappend the unique homes for every unique combination of coordinates
		str := strconv.Itoa(int(home[0])) + "," + strconv.Itoa(int(home[1]))

		if !unique[str] {
			result = append(result, home)
			unique[str] = true
		}
	}

	return result
}
