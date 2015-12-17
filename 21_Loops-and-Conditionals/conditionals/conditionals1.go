package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		rem3 := i % 3
		rem5 := i % 5
		switch {
		case rem3 == 0 && rem5 == 0:
			fmt.Print("FizzBuzz ")
		case rem3 == 0 && rem5 != 0:
			fmt.Print("Fizz ")
		case rem3 != 0 && rem5 == 0:
			fmt.Print("Buzz ")
		case rem3 != 0 && rem5 != 0:
			fmt.Print(i, " ")
		}
	}
}