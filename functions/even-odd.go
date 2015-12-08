package main

import "fmt"

func evenOdd(num int) (float64, bool) {
		return float64(num)/2, num%2 == 0
}

func main() {
	n, even := evenOdd(23)
	fmt.Println(n, even)
}
