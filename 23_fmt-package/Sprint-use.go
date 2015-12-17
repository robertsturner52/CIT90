package main

import "fmt"

func main() {
	slice := []int{4, 3, 2, 1}
	printSlice(slice)
}

func printSlice(input []int) {
	fmt.Println("first - " + fmt.Sprint(input[0]))
	fmt.Println("second - " + fmt.Sprint(input[1]))
	fmt.Println("third - " + fmt.Sprint(input[2]))
	fmt.Println("fourth - " + fmt.Sprint(input[3]))
}
