package main

import (
	"fmt"
	"math"
)

func main(){
	var up float64
	fmt.Println("Enter a decimal number:")
	fmt.Scan(&up)
	fmt.Println(math.Ceil(up))
}
