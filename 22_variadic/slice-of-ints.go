package main

import "fmt"

func main() {
	slice := []int{4, 3, 2, 1}
	variad(slice...)
	variad(4, 3, 2, 1)
}

func variad(input ...int) {
	fmt.Println(input)
}
