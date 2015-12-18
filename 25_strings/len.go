package main

import "fmt"

func main() {
	var message = "This is a string literal,\non two line(s)."
	fmt.Println(message)
	fmt.Println(len(message))
	fmt.Println("This is a string literal, on one line(s).")
	fmt.Println(len("This is a string literal, on one line(s)."))
}
