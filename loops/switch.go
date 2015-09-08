package main

import "fmt"

func main() {
	switch "right" {
	case "incorrect":
		fmt.Println("Wrong!")
		fallthrough
	case "right":
		fmt.Println("Absolutely right!")
		fallthrough
	case "correct":
		fmt.Println("Again, absolutely right!")
		fallthrough
	default:
		fmt.Println("nothing")
	}
}
