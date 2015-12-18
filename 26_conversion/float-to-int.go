package main

import (
	"fmt"
)

func main() {
	s1 := 50.5
	s2 := 24.5
	s3 := 34.0
	s4 := 22.0
	s5 := 33.5
	fmt.Println((s1 + s2 + s3 + s4 + s5)/10)
	fmt.Println(int((s1 + s2 + s3 + s4 + s5)/10)+1)
}