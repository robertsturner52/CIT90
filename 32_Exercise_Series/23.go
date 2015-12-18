package main

import "fmt"

func main() {
	slints := []int{1, 2, 3, 5, 8, 13, 21, 34}
	slintNew := []int{45, 79}
	slints = append(slints, slintNew...)
	fmt.Println(slints)
}

