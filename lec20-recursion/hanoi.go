package lec20recursion

import "fmt"

//Hanoi(3,1,3)
func Hanoi(n, from, to int){
	if n > 0 {
		Hanoi(n - 1, from, 6 - from - to)
		fmt.Printf("move from %v to %v\n", from, to)
		Hanoi(n - 1, 6 - from - to, to)
	}
}