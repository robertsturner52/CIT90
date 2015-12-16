package main

import "fmt"

func findRemainder(x int, y int) int {
	remainder := x % y
	return remainder
}

func main() {
	var inp1, inp2 int
	fmt.Println("Enter an integer: ")
	fmt.Scan(&inp1)
	fmt.Println("Enter another integer: ")
	fmt.Scan(&inp2)
	fmt.Println("The remainder of", inp1, "divided by", inp2, "is", findRemainder(inp1, inp2))
}