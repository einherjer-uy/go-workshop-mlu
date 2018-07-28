package main

import (
	"fmt"
	"errors"
	"time"
)

func main2() {
	go panicFunc2()
	go panicFunc2()
	go panicFunc2()
	go panicFunc2()

	time.Sleep(2*time.Second)
	fmt.Println("Program ends gracefully")
}

func panicFunc2(){
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic recovered. Error:%v\n",err)
		}
	}()

	panic(errors.New("forced panic"))
}

// START OMIT
func main() {
    for i:=0; i<5; i++ {
		panicFunc(i)
    }
    fmt.Println("Program ends gracefully")
}

func panicFunc(x int){
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic recovered. Error was: %v\n",err)
		}
	}()
	if x == 3 {
		panic(errors.New("pannnnic"))	
	}
	fmt.Printf("%d OK \n", x)
}
// END OMIT

