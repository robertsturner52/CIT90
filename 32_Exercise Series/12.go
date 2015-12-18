package main

import "fmt"

func main() {
	lincoln1 := "Four score and seven years ago..."
	lincoln2 := ", our fathers brought forth on this continent"
	fmt.Println(lincoln1[:len(lincoln1)-3]+lincoln2+lincoln1[len(lincoln1)-3:len(lincoln1)])
}

