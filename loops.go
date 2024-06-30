package main

import "fmt"

func loops() {
	// for loop in full form contains 3 components separated by semicolons
	/// the init statement: executed before the first iteration
	/// the condition expression: evaluated before every iteration
	/// the post statement: executed at the end of every iteration
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// but the init statement and post statement is optional (you can drop to become a while loop)
	// a while loop in C or Java is for loop in Go
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println("sum:", sum)
}