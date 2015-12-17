package main

import "fmt"

func main() {
	slice := []string{"Bob", "Todd", "Jim"}
	variad(slice...)
	variad("Bob", "Todd", "Jim")
}

func variad(input ...string) {
	fmt.Println(input)
}