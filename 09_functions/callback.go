package main

import "fmt"

func average(numbers []float64, callfunc func(float64)) {
	for _, n := range numbers {
		callfunc(n)
	}
}

func main() {
	average([]float64{23.4, 74, 87.334, 32}, func(n float64) {
		fmt.Println(n)
	})
}



