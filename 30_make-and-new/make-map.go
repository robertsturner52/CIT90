package main

import "fmt"

func main() {
	var obj map[int]string = make(map[int]string)
	fmt.Println(obj)
	fmt.Println(obj[0])
	obj[0] = "Bob"
	fmt.Println(obj[0])
}
