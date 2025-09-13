package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Create an empty integer slice with initial capacity of 3
	numbers := make([]int, 0, 3)
	
	fmt.Println("Enter integers to add to the slice (or 'X' to quit):")
	
	for {
		var input string
		fmt.Print("Enter an integer: ")
		fmt.Scanln(&input)
		
		// Check if user wants to quit
		if strings.ToUpper(input) == "X" {
			fmt.Println("Exiting...")
			break
		}
		
		// Convert string to integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid integer or 'X' to quit.")
			continue
		}
		
		// Add integer to slice
		numbers = append(numbers, num)
		
		// Sort the slice
		sort.Ints(numbers)
		
		// Print the sorted slice
		fmt.Printf("Sorted slice: %v\n", numbers)
		fmt.Println()
	}
}