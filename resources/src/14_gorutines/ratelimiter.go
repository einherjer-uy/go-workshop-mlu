package main
import "time"
import "fmt"

func main() {

	// START1 OMIT
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
	close(requests)
	// STOP1 OMIT

	// START2 OMIT
    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
	}
	// STOP2 OMIT
}