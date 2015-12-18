package main

import "fmt"

func main() {
	var def []string = make([]string, 50, 100)
	fmt.Printf("%T\n", def)
	fmt.Println(len(def))
	fmt.Println(cap(def))
}


