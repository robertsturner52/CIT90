package main

import "fmt"

func main() {
	var i int
	fmt.Println("The number is", i)
	fmt.Println("Type a number:")
	fmt.Scan(&i)
	fmt.Println("The number is now", i)
}

