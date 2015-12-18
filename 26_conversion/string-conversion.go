package main

import (
	"fmt"
	"strconv"
)

func main() {
	birth, _ := strconv.Atoi("1980")
	age := 2015 - birth
	fmt.Println("Her age is " + strconv.Itoa(age) +".")
}

