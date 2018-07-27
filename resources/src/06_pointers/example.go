package main

import (
	"fmt"
)

func main() {
	a := 10

	increment(a)
	fmt.Printf("Pasando por valor -> value: %d \n", a)
	incrementPointer(&a)
	fmt.Printf("Pasando por referencia -> value: %d \n", a)
}

func increment(a int) {
	a++
}

func incrementPointer(a *int) {
	*a++
}