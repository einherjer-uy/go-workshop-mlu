package main

import "fmt"

func genFib () func() float64 {
	var a float64 = 0
	var b float64 = 1

	return func() float64 {
		a,b = b,a+b
		return a;
	}
}

func main() {
	fib := genFib()

	for i := 0; i < 10; i++ {
		num := fib()
		fmt.Printf("%.f \n", num)
	}
}
