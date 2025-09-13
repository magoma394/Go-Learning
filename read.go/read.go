package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Name struct with fname and lname fields
// Each field will hold a string (up to 20 characters as specified)
type Name struct {
	fname string
	lname string
}

func main() {
	// Prompt user for file name
	fmt.Print("Enter the name of the text file: ")
	var fileName string
	fmt.Scanln(&fileName)

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a slice to store Name structs
	var names []Name

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by space to get first and last name
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			// Create a new Name struct
			name := Name{
				fname: parts[0],
				lname: parts[1],
			}

			// Add the struct to the slice
			names = append(names, name)
			fmt.Printf("Added: %s %s\n", name.fname, name.lname)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Print all names from the slice
	fmt.Printf("\nTotal names read: %d\n", len(names))
	fmt.Println("\nNames read from file:")
	for i, name := range names {
		fmt.Printf("%d. First Name: %s, Last Name: %s\n", i+1, name.fname, name.lname)
	}
}
