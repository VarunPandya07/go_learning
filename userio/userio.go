package main

import (
	"fmt"
)

func main() {
	fmt.Println("what's your names.....?")

	var name string

	fmt.Scanln(&name)
	fmt.Println("hello mr.", name)
}
