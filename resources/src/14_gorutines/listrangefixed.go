package main

import (
	"fmt"
	"time"
)

func main(){
	nums := [...]int{1,2,3,4,5,6,7,8,9}

	for i, n := range nums{
		go func(i, n int){
			nums[i] = n * 2
		}(i, n)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(nums)
}