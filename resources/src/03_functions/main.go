package main

import "fmt"

func sumWithLimit(x, y, limit int) (sum int, err error) {
	sum = x + y
	if sum > limit {
		err = fmt.Errorf("limit exceeded")
	}
	return
}

func main() {
	sum, err := sumWithLimit(1, 2, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(sum)	
	}
}
