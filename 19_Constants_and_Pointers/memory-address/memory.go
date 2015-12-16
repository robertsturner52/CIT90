package main

import "fmt"

const n = 42
var x = n + 1

func main() {
	fmt.Println("Constant n =", n, "and I cannot take the address of", n, ".")
	fmt.Println("Variable x =", x, "which is at", &x, ".")
}
