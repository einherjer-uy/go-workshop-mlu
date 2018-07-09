package main

import (
	"fmt"
)

func main() {
	a := [2]string{"1", "2"}
	b := [...]string{"1", "2"}
	c := [2]string{"1","3"}

	fmt.Println(c[0], c[len(c)-1] )

	for i, value := range a {
        fmt.Printf("%d:%s\n", i, value)
    }

	fmt.Println( a == b, a == c)
}
