package main

import (
	"fmt"
	"time"
)

func main(){
	nums := [...]int{1,2,3,4,5,6,7,8,9}

	for i, _ := range nums{
		go func(i int){
			nums[i] *= 2
		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(nums)
}