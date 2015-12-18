package main

import "fmt"

func main() {
	var zip *int = new(int)
	fmt.Println(zip)
	fmt.Println(*zip)
}
