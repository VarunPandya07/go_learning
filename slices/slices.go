package main
import (
	"fmt"
)
func main() {
	fmt.Println("I learn slices and make in go lang")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	numbers = append(numbers, 12, 35, 25, 11, 222)
	fmt.Println("Numbers:", numbers)

	fmt.Printf("the type of number is %T", numbers)
	fmt.Println("lenth about slice is: ", len(numbers))

	// empty slices
	name := []string{}
	fmt.Println("strings: ", name)

	// without make function
	persion := []int{}
	fmt.Println("slice:", persion)
	fmt.Println("Length:", len(persion))
	fmt.Println("capacity:", cap(persion))

	// with make function
	cars := make([]int, 3, 5)
	cars = append(cars, 25, 225, 20)
	fmt.Println("slice:", cars)
	fmt.Println("Length:", len(cars))
	fmt.Println("capacity:", cap(cars))

}
