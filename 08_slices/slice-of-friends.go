package main
import "fmt"

func main() {
	sliceOfFriends := []string{"Elijjah", "Isaiah", "Jonathan", "Mikey", "Bruce", "Phou",}
	fmt.Println(sliceOfFriends)
	fmt.Println(sliceOfFriends[1:5])
	fmt.Println(sliceOfFriends[1:])
	fmt.Println(sliceOfFriends[:5])
	newSlice := []string{"Jim", "Bob", "Frank"}
	sliceOfFriends = append(sliceOfFriends, newSlice...)
	fmt.Println(sliceOfFriends[1:])
}
