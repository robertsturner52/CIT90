package main
import "fmt"

func main() {
	billion := 1000000000
	fmt.Printf("%T\n", billion)
	fmt.Println(billion)
	fmt.Printf("%d\t %b\t %#x\n", billion, billion, billion)
	ma := &billion
	fmt.Println(ma)
	fmt.Printf("%T\n", ma)
	fmt.Printf("%d\t %b\t %#x\n", ma, ma, ma)
}
