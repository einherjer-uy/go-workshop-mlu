package main

import (
	"sync"
	"fmt"
)

func main() {
    counter := new(int)
	var wg sync.WaitGroup
	
    // multiple producers
    for producerId := 0; producerId < 1000; producerId++ {
        wg.Add(1)

        // spawn a new counter
        go func(producerId int) {
            defer wg.Done()
            *counter = *counter + 1
        }(producerId)
    }

    // wait for the N producers
    wg.Wait()

    // check expected value
    fmt.Println(*counter)
}