package main

import (
	"fmt"
)

func charFreq(word string) map[string]int {
 	freq := make(map[string]int)
	
	for _, l := range word {
		freq[string(l)]++
	}
	
	return freq
}

func concurrentCharFreq(words []string) map[string]int{
	result := make(map[string]int)
	c := make(chan map[string]int, len(words))
	
	for _, word := range words{
		go func(word string){
			c <- charFreq(word)
		}(word)
	}
	
	for i:=0; i<len(words); i++ {
		freq := <- c
		for k,v := range freq {
			result[k] += v
		}
	}
	
	return result
}

// START OMIT
func main(){
	fmt.Println(concurrentCharFreq([]string{"hola", "chau", "chorizo"}))
}
// END OMIT