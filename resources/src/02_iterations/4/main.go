package main

import (
	"fmt"
	"runtime"
)

func main() {
    switch runtime.GOOS {
    case "darwin":
        fmt.Println("Go runs on OS X.")
    case "linux":
        fmt.Println("Go runs on Linux.")
    default:
        fmt.Printf("Go runs on %s.", runtime.GOOS)
    }
}
