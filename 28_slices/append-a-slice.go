package main

import "fmt"

func main() {
	the1stSlice := []string{"a","ab","abc","abcd"}
	the2ndSlice := []string{"a","ab","abc","abcd"}
	the1stSlice = append(the1stSlice, the2ndSlice...)
	fmt.Println(the1stSlice)
}
