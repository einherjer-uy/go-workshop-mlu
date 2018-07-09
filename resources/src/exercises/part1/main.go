package main

import (
	"fmt"
)

func multiples(to int) []int {
	var result []int
	for i := 0; i <= to; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, i)
		}
	}
	return result
}

// START OMIT
func main() {
	fmt.Printf("%v", multiples(100))
}
// END OMIT