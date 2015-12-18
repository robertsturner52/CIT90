package main

import "fmt"

func main() {
	var char rune
	char = 'r'
	fmt.Println(char)
	fmt.Printf("%T\n", char)
	fmt.Println(string(char))
}

