package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		if (i % 3 == 0) {
			sum = sum + i
			fmt.Print(i, " ")
		} else {
			if (i % 5 == 0) {
				sum = sum + i
				fmt.Print(i, " ")
			}
		}
	}
	fmt.Print("The sum of all the multiples of 3 or 5 below 1000 is ", sum, ".")
}
