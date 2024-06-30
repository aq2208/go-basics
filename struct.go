package main

import "fmt"

type Car struct {
	name  string
	color string
	price int
}

func structDemo() {
	var car = Car{"BMW", "white", 2}
	fmt.Println(Car{"BMW", "white", 2})
	fmt.Println("Car name: ", car.name)
}
