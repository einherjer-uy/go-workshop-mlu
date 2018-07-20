package main

import (
	"fmt"
	"time"
)

func main(){
	nums := [...]int{1,2,3,4,5,6,7,8,9}

	for _, n := range nums{
		go func(){
			n = n * 2
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println(nums)
}