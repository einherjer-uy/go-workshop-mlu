package main

import "fmt"

func split (arr []int, condition func(i int) bool) (a, b []int) {
	for _, e := range arr {
		if(condition(e)){
			a = append(a, e)
		}else{
			b = append(b, e)
		}
	}
	return
}

func main() {
	nums := []int{1,2,3,4,5}
	
	f := func(x int) bool {
		return x % 2 == 0
	}

	pares, impares := split(nums, f)

	fmt.Println(pares)
	fmt.Println(impares)
	
}
