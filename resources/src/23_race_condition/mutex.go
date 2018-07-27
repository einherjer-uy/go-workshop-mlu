package main

import (
	"sync"
	"fmt"
)

func main() {
    var counter int
    var mu sync.Mutex // HL
    var wg sync.WaitGroup

    for producerId := 0; producerId < 1000; producerId++ {
        wg.Add(1)
        // spawn a new counter
        go func(producerId int) {
            mu.Lock() // HL
            defer mu.Unlock() // HL
            defer wg.Done()
            counter = counter + 1
        }(producerId)
    }
    wg.Wait()           // wait for the N producers
    fmt.Println(counter)
}