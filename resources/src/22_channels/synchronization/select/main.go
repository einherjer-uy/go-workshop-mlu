package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
    work := []int{1, 2, 3, 4, 5}
    out := make(chan int)

    for _, w := range work {
        go func(w int) {
            time.Sleep(time.Duration(w*100) * time.Millisecond)
            out <- w*10
        }(w)
    }

	timeout := time.After(300 * time.Millisecond)
    for i:=0; i<len(work); i++ {
        select {
        case v := <- out:
            fmt.Printf("Received %d \n", v)
        case <-timeout:
            fmt.Print("Timeout")
            return
        }
    }
    fmt.Print("Done")
}
// END OMIT