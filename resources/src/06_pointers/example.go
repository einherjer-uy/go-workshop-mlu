package main

import (
	"fmt"
)

func main() {
	a := 10

	increment(a)
	fmt.Printf("After increment method -> value: %d \n", a)
	incrementPointer(&a)
	fmt.Printf("After incrementPointer method -> value: %d \n", a)
}

func increment(a int) {
	a++
}

func incrementPointer(a *int) {
	*a++
}