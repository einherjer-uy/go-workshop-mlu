package main

import (
	"sync"
	"fmt"
	"sync/atomic"
)

func main() {
    counter := new(int32) // HL
    var wg sync.WaitGroup
    // multiple producers
    for producerId := 0; producerId < 1000; producerId++ {
        wg.Add(1)

        // spawn a new counter
        go func(producerId int) {
            defer wg.Done()
            atomic.AddInt32(counter, 1) // HL
        }(producerId)
    }
    wg.Wait()                            // wait for the N producers
    fmt.Println(atomic.LoadInt32(counter)) // HL
}