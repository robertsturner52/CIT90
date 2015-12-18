package main

import "fmt"

func main() {
	word := make(map[string]string)
	word["dog"] = "loyal pet"
	word["cat"] = "aloof pet"
	word["pig"] = "smart pet"
	fmt.Println(word["cat"])
	word["cat"] = "soft pet"
	fmt.Println(word["cat"])
}


