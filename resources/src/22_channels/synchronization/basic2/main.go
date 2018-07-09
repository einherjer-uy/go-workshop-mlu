package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	work := []int{1, 2, 3, 4, 5}
	out := make(chan int)

	for w := range work {
		go func(w int) {
			time.Sleep(time.Second)
			out <- w*10
		}(w)
	}
	for i:=0; i<len(work); i++ {
		fmt.Printf("Received %d \n", <-out)
	}
	fmt.Print("Done")
}
// END OMIT