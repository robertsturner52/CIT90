package main

import "fmt"

type person struct {
	name	string
	city	string
	age		int
}

func main() {
	per01 := person{"Bob", "Clovis", 63}
	fmt.Println(per01.name)
	fmt.Println(per01.city)
	fmt.Println(per01.age)
}


