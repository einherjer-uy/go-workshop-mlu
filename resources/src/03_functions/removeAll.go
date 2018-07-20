package main

import "fmt"

func removeAll (arr []int, shouldRemove func(i int) bool) []int {
	var filteredArr []int
	
	for _, e := range arr {
		if(!shouldRemove(e)){
			filteredArr = append(filteredArr, e)
		}
	}

	return filteredArr
}

func main() {
	nums := []int{1,2,3,4,5}
	f := func(x int) bool {
		return x % 2 != 0
	}

	fmt.Println(removeAll(nums, f))
}
