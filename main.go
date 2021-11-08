package main

import "fmt"

func main() {
	hash_map := NewHashMap()
	hash_map.Push("first", 67)

	fmt.Println(hash_map.elements)
	if value, err := hash_map.Get("first"); err == nil {
		fmt.Println(value)
	}
}
