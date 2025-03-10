package main

import "fmt"

func main() {
	fmt.Println("i learning about if else condition for go-lang")
	var num1, num2 int

	fmt.Println("enter num1:")
	fmt.Scanln(&num1)
	fmt.Println("enter num2:")
	fmt.Scanln(&num2)
	if num1 > num2 {
		fmt.Println("num1 is gretner then num2")
	} else if num1 < num2 {
		fmt.Println("num2 is greater then num1")
	} else if num1 == num2 {
		fmt.Println("both numbers are equal")
	} else {
		fmt.Println("enter the velid numbers")
	}
}
