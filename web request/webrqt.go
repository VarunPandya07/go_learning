package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learning about the web request")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET responce", err)
		return
	}
	defer res.Body.Close()

	// fmt.Println("Type of responce: %T\n", res)
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("not data is found", err)
		return
	}
	fmt.Println("responce:", string(data))
}
