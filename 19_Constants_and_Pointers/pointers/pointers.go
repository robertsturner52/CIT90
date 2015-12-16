package main

import "fmt"

func zero(x int) int {
	x = 0
	return x
}

func point(x *int) *int {
	*x = 0
	return x
}

func main() {
	x := 5
	zero(x)
	fmt.Println(zero(x))
	fmt.Println(x)
	fmt.Println(point(&x))
	fmt.Println(x)
}
