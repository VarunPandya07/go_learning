package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"iserId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("learing GET request in go")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error getting:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("error getting response", res.Status)
	}
	/*
	   its genric response

	   data, err := ioutil.ReadAll(res.Body)

	   	if err != nil {
	   		fmt.Println("Error reading:", err)
	   		return
	   	}

	   fmt.Println("data is:", string(data))
	*/

	/*unmarshalling code*/
	var todo Todo

	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error decoding..!", err)
		return
	}

	fmt.Println("Todo:", todo)
}
