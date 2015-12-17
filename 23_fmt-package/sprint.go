package main

import "fmt"

func main() {
	slice := []int{4, 3, 2, 1}
	printSlice(slice)
}

func printSlice(input []int) {
	fmt.Println(input)
	fmt.Printf("%T", input)
	fmt.Println()
	fmt.Println(fmt.Sprint(input))
	fmt.Printf("%T", fmt.Sprintln(input))
}
