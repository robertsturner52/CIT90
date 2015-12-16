package main

import "fmt"

const (
	_ = iota
	counter1 = iota
	counter2 = iota
)

func main() {
	fmt.Println(counter1 * 100, counter1 * 200, counter1 * 300)
	fmt.Println(counter2 * 100, counter2 * 200, counter2 * 300)
}