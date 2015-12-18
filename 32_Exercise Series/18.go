package main

import "fmt"

func main() {
	var indef interface{}
	fmt.Println(indef)
	fmt.Printf("%T\n", indef)
	indef = "laughing cow"
	fmt.Println(indef)
	fmt.Printf("%T\n", indef)
	str := indef.(string)
	fmt.Println(str)
	fmt.Printf("%T\n", str)
}
