package main

import "fmt"

func main() {
	fmt.Println("i learn switch cases in go-lang")

	day := 15
	switch day {
	case 12:
		fmt.Println("monday")
	case 13:
		fmt.Println("tuseday")
	case 14:
		fmt.Println("wenrsday")
	case 15:
		fmt.Println("tersday")
	case 16:
		fmt.Println("friday")
	default:
		fmt.Println("etner velid number")
	}
}
