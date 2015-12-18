package main

import "fmt"

func main() {
	var truth *bool = new(bool)
	fmt.Println(truth)
	fmt.Println(*truth)
}
