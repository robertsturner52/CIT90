package main

import "fmt"

func main() {
	var city string
	fmt.Print("Enter your home city: ")
	fmt.Scan(&city)
	fmt.Println("My home is in " + city + ".")
}
