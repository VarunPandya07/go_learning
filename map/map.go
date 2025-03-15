package main

import (
	"fmt"
)

func main() {
	fmt.Println("learning about map")

	StudentGrades := make(map[string]int)

	StudentGrades["varun"] = 100
	StudentGrades["harsh"] = 85
	StudentGrades["daksh"] = 58
	StudentGrades["yash"] = 85

	fmt.Println("varun grades ", StudentGrades["varun"])
	StudentGrades["daksh"] = 0
	fmt.Println("daksh grades ", StudentGrades["daksh"])

	delete(StudentGrades, "harsh")
	fmt.Println("harsh grades ", StudentGrades["harsh"])

	// checking existing the key
	grade, exists := StudentGrades["rakesh"]
	fmt.Println("grade of rakesh is:", grade)
	fmt.Println("exists of rakesh is", exists)

	// using for loop in the map
	for index, value := range StudentGrades {
		fmt.Printf("key is %s, and value is %d\n", index, value)
	}
}
