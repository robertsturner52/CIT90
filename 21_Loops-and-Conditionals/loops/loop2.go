package main

import "fmt"

func main() {
	i := 1
	for {
		j := i % 5
		fmt.Print(j, " ")
		if i > 10 {
			break
		}
		i = i + 1
	}
}
