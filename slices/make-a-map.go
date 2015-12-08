package main
import "fmt"
func main() {
	myMap := map[string]string{
		"sequoia": "seven-letter word with all five vowels",
		"rodieau": "made-up word with all five vowels",
	}
	myMap["newWord"] = "This is a new word."
	fmt.Println(myMap)
	fmt.Println(myMap["newWord"])
	myMap["newWord"] = "This is the changed word."
	fmt.Println(myMap["newWord"])
	delete(myMap, "newWord")
	for key, val := range myMap {
		fmt.Println(key, " - ", val)
	}
	fmt.Println(len(myMap))
	if _, exists := myMap["newWord"]; exists {
		fmt.Println("This map has a new word.")
	} else {
		fmt.Println("This map has no new word.")
	}
}