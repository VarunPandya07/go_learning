package main

import (
	"fmt"
)

func simplefunction() {
	fmt.Println("this is my simple function in golang")
	fmt.Println("-------------------------------------")
	fmt.Println()
}

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("i learning about the function in golang")
	fmt.Println("-------------------------------------")
	fmt.Println()

	simplefunction()
	sum := add(5, 10)
	fmt.Println("addition of this two num is:", sum)
}
