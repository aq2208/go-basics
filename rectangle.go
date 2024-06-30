package main

import (
	"fmt"
)

func area(length int, width int) int {
	return length * width
}

func perimeter(length int, width int) int {
	return (length + width) * 2
}

func rectangle() {
	// declare variables for rectangle's length and width
	var length, width int

	// get user input for rectangle's length and width
	fmt.Println("Input length: ")
	fmt.Scan(&length)
	fmt.Println("Input width: ")
	fmt.Scan(&width)

	// calculate the area and perimeter
	fmt.Printf("Rectangle perimeter: %v\n", perimeter(length, width))
	fmt.Printf("Rectangle area: %v\n", area(length, width))
}