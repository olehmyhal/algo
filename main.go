package main

import (
	"fmt"
)

func main(){
	fmt.Print("Hello world\n")

	Arangements([]int{}, 3, 2)
}

func Arangements(prefix []int, n, k int){
	if len(prefix) == k {
		fmt.Printf("%v\n", prefix)
		return 
	}

	for c := 1; c < n + 1; c++ {
		if len(prefix) == 0 || c > prefix[len(prefix) - 1] {
			Arangements(append(prefix, c), n, k)
		}
	}
}