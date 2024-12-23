package lec11sort

import "fmt"

//LoopMergeSort([]int{1, -10, -100, 11, 1000, 0 ,12})
func mergeLoop(l,r,m int, arr []int){
	aux := []int{}

	for i := l; i < m; i++ {
		aux = append(aux, arr[i])
	}

	fmt.Printf("AUX %v \n", aux)

	i, j, k := 0, m, l

	for i < len(aux) && j < r {
		currI, currJ := aux[i], arr[j]
		if currI < currJ {
			arr[k] = currI
			i++
		} else {
			arr[k] = currJ
			j ++
		}
		k++
	}

	for i < len(aux) {
		arr[k] = aux[i]
		i++ 
		k++
	}

	for j < r {
		arr[k] = arr[j]
		j++ 
		k++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func LoopMergeSort(nums []int){
	n := len(nums)

	for k := 1; k < n; k *= 2 {
		fmt.Printf("k %v nums are %v \n", k, nums)
		for i := 0; i < n; i += 2 * k {
			l, r := i, min(n, i + 2*k)

			m := min(i+k, n)

			fmt.Printf("l %v r %v m %v\n", l, r,m)

			mergeLoop(l, r, m, nums)
		}
	}

	fmt.Printf("%v", nums)
}