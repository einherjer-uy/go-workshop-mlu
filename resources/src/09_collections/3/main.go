package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int) 
	var m2 map[int]int 

	m[1] = 1
	m[2] = 2
	m[3] = 3
	delete(m,3)

	fmt.Println(m2 == nil) //panic if m2[0]=1 (assignment to nil map)

	const KEY = 2
	value, exists := m[KEY]
	fmt.Printf("Key: %d, Value: %d, Exists: %t", KEY, value, exists)

}
