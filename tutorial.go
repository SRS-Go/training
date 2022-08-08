package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 7,
	}
	fmt.Println(mp)
	fmt.Println(mp["orange"]) //access the element
	mp["apple"] = 20          // update the value
	fmt.Println(mp)
	mp["banana"] = 3 // add new key value pair
	fmt.Println(mp)
	delete(mp, "pear") //delete an element
	fmt.Println(mp)
	mp1 := make(map[string]int) // define empty map
	mp1["good"] = 1
	fmt.Println(mp1)

	// check if key exists
	val, ok := mp["apples"]
	fmt.Println(val, ok)
}
