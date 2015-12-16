package main

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
)

func main() {
	u, err := uuid.ParseHex("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(u)
}
