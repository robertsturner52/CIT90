package main

import "fmt"

func main() {
	pet := make(map[string]string)
	pet["dog"] = "Fido"
	pet["cat"] = "Tiger"
	pet["pig"] = "Arnold"
	fmt.Println(pet)
	delete(pet, "cat")
	pet["cat"] = "Shadow"
	pet["dog"] = "Sparky"
	for key, val := range pet {
		fmt.Println(key, " - ", val)
	}
	fmt.Println(len(pet))
	if val, exists := pet["pig"]; exists {
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
		pet["pig"] = "Frank"
		fmt.Println("\"pig\" exists")
		fmt.Println("His name is ", pet["pig"])
	} else {
		fmt.Println("There is no \"pig\" in the poke.")
	}
	delete(pet, "pig")
	if val, exists := pet["pig"]; exists {
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
		pet["pig"] = "Frank"
		fmt.Println("\"pig\" exists")
		fmt.Println("His name is now ", pet["pig"])
	} else {
		fmt.Println("There is no \"pig\" in the poke.")
	}
}
