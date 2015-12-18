package main

import (
	"fmt"
)

func main() {
	fmt.Println(string([]byte{'s','r','i','r','a','c','h','a',' ','s','a','u','c','e'}))
	fmt.Println(string([]rune{115, 114, 105, 114, 97, 99, 104, 97, 32, 115, 97, 117, 99, 101}))
}