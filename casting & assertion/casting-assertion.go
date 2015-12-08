package main

import "fmt"

func main() {
	carSlice := []string{"Ford", "Chevrolet", "Chrysler"}
	carSlice = append(carSlice, "Kia", "Nissan")
	fmt.Println(carSlice)
	newCarSlice := []string{"Honda", "Tesla", "BMW"}
	carSlice = append(carSlice, newCarSlice...)
	fmt.Println(carSlice)

	trainSlice := make([]string, 3, 6)
	trainSlice[2] = "BNSF"
	trainSlice[1] = "Southern Pacific"
	fmt.Println(len(trainSlice))
	fmt.Println(trainSlice[0])

}
