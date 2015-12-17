package main

import "fmt"

func isEven(x int) bool {
	even := x & 0x01 == 0
	return even
}

func main() {
	var input int
	var str string
	var strinput float64
	fmt.Println("Enter an integer: ")
	fmt.Scan(&input)
	fmt.Println(isEven(input))
	fmt.Println("Enter an number: ")
	fmt.Sscan(str, strinput)
	fmt.Println(str)
	fmt.Println(strinput)
	fmt.Printf("%T", strinput)
	fmt.Println()
	fmt.Printf("%T", str)
}
