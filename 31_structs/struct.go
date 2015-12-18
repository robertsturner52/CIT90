package main

import "fmt"

type info struct {
	name	string
	city	string
	age		int
}

func main() {
	info01 := info{"Bob", "Clovis", 63}
	info02 := info{"Jim", "Fresno", 45}
	info03 := info{"JimBob", "Fresno", 21}
	fmt.Println(info01)
	fmt.Println(info01.name)
	fmt.Println(info02.name)
	fmt.Println(info03.name)
	fmt.Println(info03.city)
	info03.city = "Minkler"
	fmt.Println(info03.city)
}

