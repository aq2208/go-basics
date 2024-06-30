package main

import (
	"fmt"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func swap(x int, y int) (int, int) {
	return y, x
}

// function main -> run first at compile time
func main2() {
	fmt.Println("Hello world!")
	fmt.Printf("Random number: %v\n", rand.Int())
	fmt.Println(add(10, 20))
	fmt.Println(swap(1, 2))
}