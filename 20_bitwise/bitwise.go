package main

import "fmt"

func isEven(x int) bool {
	even := x & 0x01 == 0
	return even
}

func main() {
	var input int
	fmt.Println("Enter an integer: ")
	fmt.Scan(&input)
	fmt.Println(isEven(input))
}
