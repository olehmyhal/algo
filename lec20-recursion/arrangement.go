package lec20recursion

import "fmt"

//Arrangement([]int{}, 3)
func Arrangement(prefix []int, n int){
	if len(prefix) == n {
		fmt.Printf("%v\n", prefix) 
		return
	}

	Arrangement(append(prefix, 0), n)
	Arrangement(append(prefix, 1), n)
}