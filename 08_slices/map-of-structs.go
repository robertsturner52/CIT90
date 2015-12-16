package main
import "fmt"
func main() {
	myMap := map[string]string{
		"sequoia": "seven-letter word with all five vowels",
		"rodieau": "made-up word with all five vowels",
	}
	fmt.Println(myMap)
	fmt.Println(myMap["sequoia"], myMap["rodieau"])
}