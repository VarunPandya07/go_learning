package main

import "fmt"

func main() {
	fmt.Println("i larning array in go lang")

	var name [5]string
	name[0] = "varun"
	name[2] = "pandya"
	fmt.Println("name of persion", name)

	var number = [5]int{1, 2, 3, 4, 5}
	fmt.Println("number is:", number)

	fmt.Println("lanth of array is: ", len(number))

	fmt.Println("value of set in array no. of 2 is:", name[2])

	// empty array

	var price [5]int
	fmt.Println("price is:", price)

	var persion [5]string
	fmt.Printf("price is: %q", persion)

}
