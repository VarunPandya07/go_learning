package main

import (
	"fmt"
)

func main() {
	fmt.Println("learning about for loop")

	// for loop using in normal way
	for i := 0; i <= 4; i++ {
		fmt.Println("Number is:", i)
	}

	// range key word use in for loop
	number := []int{11, 12, 73, 54, 25, 36}
	for index, value := range number {
		fmt.Printf("index is %d, value is %d\n", index, value)

	}

	// for loop using in string 
	data := "Hello world!"

	for index, char := range data{
		fmt.Printf("index number is %d, charater is %c\n", index, char)
	}
}
