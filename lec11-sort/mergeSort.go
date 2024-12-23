package lec11sort

import "fmt"

//MergeSort([]int{-10, -100, 1000, 1, 0 ,12})

func sort(l,r int, arr []int){
	if r - l < 2 {
		return
	}

	m := (r + l) / 2
	sort(l, m, arr)
	sort(m, r, arr)
	merge(l, r, m, arr)
}

func merge(l,r,m int, arr []int){
	aux := []int{}

	for i := l; i < m; i++ {
		aux = append(aux, arr[i])
	}

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

func MergeSort(nums []int){
	n := len(nums)

	sort(0, n, nums)
	
	fmt.Printf("%v", nums)
}