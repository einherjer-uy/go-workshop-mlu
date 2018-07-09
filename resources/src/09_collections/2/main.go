package main

import (
	"fmt"
)

func main() {
    a := []int{1,2,3}

    s1 := a[0:2]         // same as s1:=a[:2]
    var s2 []int         // len(s2) == 0 , s2 == nil
    s3 := make([]int, 2) // len(s3) == 2 , cap(s3) == 2. Same as make([]int,2,2)
    
    s1 = append(s1, 5)
    s2 = append(s2, 5)
    s3 = append(s3, 5)

    fmt.Println(s1 , s2 , s3)
}
