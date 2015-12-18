package main

import "fmt"

func main() {
	var r rune
	for i := 0; i < 256; i++ {
		r = rune(i)
		fmt.Println(i, "-", string(r))
	}
}

