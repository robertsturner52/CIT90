package main

import "fmt"

func main() {
	list := map[int]string{
		1: "Omaha",
		2: "Cleveland",
		3: "Dallas",
	}
	n := len(list)
	fmt.Println(n)
	for key, value := range list {
		fmt.Println("Key:", key, "Value:", value)
	}
	delete(list, 2)
	fmt.Println(len(list))
	for key, value := range list {
		fmt.Println("Key:", key, "Value:", value)
	}

}