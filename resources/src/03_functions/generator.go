package main

import "fmt"

func gen (initial int) func() int {
	a := initial - 1

	return func() int {
		a++
		return a;
	}
}

func main() {
	g1 := gen(0)
	g2 := gen(100)

	for i := 0; i <= 10; i++ {
		fmt.Printf("g1:%d \t g2:%d \n", g1(), g2())
	}
}
