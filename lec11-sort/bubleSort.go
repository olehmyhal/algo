package lec11sort

import "fmt"

//BubleSort([]int{30,10, -100, 20, 1})
func BubleSort(nums []int){
	n := len(nums)

	for j := 0; j < n; j++ {
		for i := 0; i < n - 1 - j; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}


	fmt.Printf("%v", nums)
}