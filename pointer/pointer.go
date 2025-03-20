package main

import "fmt"

func mdifyValueByRefrance(num *int) {
	*num = *num + 20
}

func main() {
	fmt.Println("learning about pointer in go")
	// var num int
	// num = 2
	num := 2
	// var ptr *int
	// ptr = &num
	ptr := &num

	// with value declartion
	fmt.Println("Num has value :", num)
	fmt.Println("pointer cantians :", ptr)

	// without value declartion
	n := 10
	p := &n

	fmt.Println("pointer cantains :", ptr)
	fmt.Println("data cantains thorugh Pointer :", *p)

	// if the pointer is nil
	var pointer *int
	if pointer == nil {
		fmt.Println("pointer is nil")
	}

	value := 10
	mdifyValueByRefrance(&value)
	fmt.Println("value cantains :", value)
}
