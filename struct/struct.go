package main

import "fmt"

type Person struct {
	FristName string
	LastName  string
	Age       int
}

type Conect struct {
	email        string
	phone_number int
}

type Address struct {
	houseNumber int
	state       string
}

type Employee struct {
	peroson_details Person
	peroson_conect  Conect
	peroson_address Address
}

func main() {
	fmt.Println("learing about struct")

	var varun Person

	fmt.Println("varun parson", varun)
	varun.FristName = "varun"
	varun.LastName = "pandya"
	varun.Age = 21
	fmt.Println("varun parson", varun)

	// sencond way to instialz the peroson struct
	parson := Person{
		FristName: "harsh",
		LastName:  "pambhar",
		Age:       19,
	}
	fmt.Println("parson is ", parson)

	// using new key word
	// its is using like pointer
	parson1 := new(Person)
	parson1.FristName = "rudesh"
	parson1.LastName = "pandya"
	parson1.Age = 25

	fmt.Println("parson1 is ", parson1)

	fmt.Println("--------------------------------------------------------------------------------------------")
	var employee1 Employee
	employee1.peroson_details = Person{
		FristName: "dev",
		LastName:  "Kotecha",
		Age:       22,
	}
	employee1.peroson_conect.email = "devk132@gmail.com"
	employee1.peroson_conect.phone_number = 1253457896

	employee1.peroson_address = Address{
		houseNumber: 14,
		state:       "gujarat",
	}

	fmt.Println("Employee 1 :", employee1)
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Println("Employee 1 :", employee1.peroson_details)
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Println("Employee 1 :", employee1.peroson_conect)
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Println("Employee 1 :", employee1.peroson_address)
	fmt.Println("--------------------------------------------------------------------------------------------")
}
