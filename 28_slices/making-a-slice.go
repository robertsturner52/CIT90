package main

import "fmt"

func main() {
	custNum := make([]int, 5, 10)
	custNum[0] = 1
	custNum[1] = 2
	custNum[2] = 3
	custNum[3] = 4
	custNum[4] = 5
	fmt.Println(custNum)
}
