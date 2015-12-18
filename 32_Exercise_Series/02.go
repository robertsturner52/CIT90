package main

import "fmt"

func main() {
	hex := fmt.Sprintf("%x", 394)
	fmt.Println(hex)
	fmt.Printf("%T\n", hex)
}