package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("learning json in go lang")
	person := Person{
		Name:    "varun",
		Age:     21,
		IsAdult: true}

	// fmt.Println("person data is: ", person)

	// converting person into JSON encoding (Marshalling)
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling", err)
		return
	}

	fmt.Println("person marshalling", string(data))

	// converting person JSON decoding (unmarshalling)
	var personData Person
	err = json.Unmarshal(data, &personData)
	if err != nil{
		fmt.Println("error unmarshalling",err)
	}
	fmt.Println("person Data is:", personData)
}
