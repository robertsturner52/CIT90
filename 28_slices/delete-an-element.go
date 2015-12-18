package main

import "fmt"

func main() {
	theSlice := []string{"a","ab","abc","abcd"}
	fmt.Println(theSlice)
	theSlice = append(theSlice[:2], theSlice[3:]...)
	fmt.Println(theSlice)
}
