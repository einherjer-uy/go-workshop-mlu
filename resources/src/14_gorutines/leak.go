package main

import (
	"fmt"
	"runtime"
)

// START3 OMIT
func main(){
	for i := 0; i < 10; i++ {
		fmt.Println(getMessage())
	}

	fmt.Println("Go routines activas:", runtime.NumGoroutine())
}
// STOP3 OMIT

// START2 OMIT
func getMessage() string{
	a := message("a")
	b := message("b")

	select{
		case adata := <-a:
			return adata
		case bdata := <-b:
			return bdata
	}
}
// STOP2 OMIT

// START1 OMIT
func message(nombre string) <-chan string{
	c := make(chan string)

	go func(){ c <- nombre }()

	return c
}
// STOP1 OMIT