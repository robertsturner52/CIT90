package main

import "fmt"

func main() {
	twoParams := func (name string, age int) (string) {
	return fmt.Sprint(name, " is ", age, " years old.")
	}
	fmt.Println(twoParams("Bob", 63))
}



