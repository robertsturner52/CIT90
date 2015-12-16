package main

import "fmt"

func main() {
	type Person struct {
        Name  string
        Likes []string
    }
    b := Person{"Bob", []string{"cheese", "crackers"}}
    fmt.Println(b)
	var people []*Person
	fmt.Println(people)

    likes := make(map[string][]*Person)
    for _, p := range people {
        for _, l := range p.Likes {
            likes[l] = append(likes[l], p)
        }
    }

    for _, p := range likes["cheese"] {
        fmt.Println(p.Name, "likes cheese.")
    }
}
