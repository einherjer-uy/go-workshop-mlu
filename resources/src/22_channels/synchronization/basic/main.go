package main

import (
	"fmt"
	"time"
	"math/rand"
)

// START OMIT
func main() {
	work := make(chan int)

	for i:=0; i<3; i++ {
		go func(id int) {
			for {
				w := <- work 
				fmt.Printf("Worker %d picked work %d \n", id, w)
			}
		}(i)
	}
	
	for i:=0; i<100; i++ {
		work <- rand.Intn(10)
		time.Sleep(time.Second)
	}
}
// END OMIT