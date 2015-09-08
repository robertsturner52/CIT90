package main
import "fmt"
func main() {
	a := "content"
	fmt.Println(a)
	fmt.Println(&a)
	var b *string = &a
	fmt.Println(b)
	fmt.Println(*b)
	*b = "new content"
	fmt.Println(a)
}
