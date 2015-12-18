package main

import "fmt"

func main() {
	slints := []int{1, 2, 3, 5, 8, 13, 21, 34}
	for i, val := range slints {
		fmt.Println(i, " - ", val)
	}
}

