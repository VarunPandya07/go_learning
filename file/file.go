package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("error while crate file:", err)
	// 	return
	// }
	// defer file.Close()

	// content := "hello pandya varun"

	// _, error := io.WriteString(file, content+"\n")
	// if error != nil {
	// 	fmt.Println("Error while writeing file", error)
	// 	return
	// }
	// fmt.Println("file was created successfully")

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("error while Open file", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("error while Reading file", err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}

}
