package main

import "fmt"

func isDividedByTwo(str string) bool {
	return len(str) % 2 == 0 
}

func stringexcercise() {
	var str string

	// get user input for rectangle's length and width
	fmt.Println("Input string: ")
	fmt.Scan(&str)

	fmt.Printf("Is string's length is divided by 2: %v\n", isDividedByTwo(str))
}