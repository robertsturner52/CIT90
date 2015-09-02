package main

import "fmt"

var name string
func main() {
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello ", name)
}
