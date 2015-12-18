package main

import "fmt"

func main() {
	var code *string = new(string)
	fmt.Println(code)
	fmt.Println(*code)
}
