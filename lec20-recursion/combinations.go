package lec20recursion

import "fmt"

//Combinations([]int{}, 3, 2)
func Combinations(prefix []int, n, k int){
	if len(prefix) == k {
		fmt.Printf("%v\n", prefix)
		return 
	}

	for c := 1; c < n + 1; c++ {
		if len(prefix) == 0 || c > prefix[len(prefix) - 1] {
			Combinations(append(prefix, c), n, k)
		}
	}
}