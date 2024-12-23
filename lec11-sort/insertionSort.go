package lec11sort

import "fmt"

//InsertionSort([]int{5,3,2,1,0,4})
func InsertionSort(nums []int){
	n := len(nums)

	for i := 0; i < n; i++ {
		Inner:
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break Inner
			}
		}
	}

	fmt.Printf("%v", nums)
}