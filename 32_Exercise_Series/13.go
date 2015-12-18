package main

import "fmt"
import "strconv"

func main() {
	num := 42
	strNum := strconv.Itoa(num)
	fmt.Println(strNum)
	fmt.Printf("%T\n", strNum)
	numStr, _ := strconv.Atoi(strNum)
	fmt.Println(numStr)
	fmt.Printf("%T\n", numStr)
}

