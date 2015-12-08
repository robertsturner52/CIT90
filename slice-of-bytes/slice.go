package main
import "fmt"

func main() {
	intro := "Four score and seven years ago...."
	fmt.Printf("%T\n", intro)
	fmt.Println(intro)
	fmt.Println([]byte(intro))
	bs := []byte(intro)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	for _, v := range bs {
	fmt.Printf("%d\t\t %b\n", v, v)
	}
}
