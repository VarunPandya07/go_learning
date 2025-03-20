package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("starting of code")
	data := add(4, 5)
	defer fmt.Println("data is:", data)
	defer fmt.Println("end of the code")
	fmt.Println("middle of the code")
}
