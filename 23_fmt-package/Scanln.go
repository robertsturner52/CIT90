package main

import "fmt"

func main() {
	var input, output int
	fmt.Println("Enter length in feet: ")
	fmt.Scan(&input)
	output = input * 12
	fmt.Println(input, "feet is", output, "inches.")
}