package main

import "fmt"

func main() {
	for i := 0; i < 1001; i++ {
		fmt.Println(i)
		i = i + 1
	}
}
