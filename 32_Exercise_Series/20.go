package main

import "fmt"

func main() {

	slints := []int{1, 2, 3, 5, 8, 13, 21, 34}
	for i := 0 ; i < len(slints); i++ {
		fmt.Println(i, " - ", slints[i])
	}
}

