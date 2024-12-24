package lec15qsort

import "fmt"

func QuickSortCaller(){
	arr := []int{4, 9, 4, 4, 8, 4, 5, 9, 4, 3}
	sortedArr := QuickSort(arr)

	fmt.Printf("unsorted: %v\n", arr)
	fmt.Printf("sorted: %v\n", sortedArr)
}

func QuickSort(arr []int)[]int{
	if len(arr) < 2 {
		return arr //Base case
	}

	pivotIndex := medianOfThree(arr)
	pivot := arr[pivotIndex]
	less, greater, eq := partion(pivot, arr)

	less = QuickSort(less)
	greater = QuickSort(greater)

	return append(append(less, eq...), greater...)
}

func partion(pivot int, arr []int)(less []int, greater []int, eq []int){
	for _, a := range arr {
		if a < pivot {
			less = append(less, a)
		} else if a == pivot {
			eq = append(eq, a)
		} else {
			greater = append(greater, a)
		}
	}

	return less, greater, eq
}

func medianOfThree(arr []int)int{
	startIndex, medianIndex, endIndex := 0, len(arr) / 2, len(arr) - 1
	start, median, end := arr[startIndex], arr[medianIndex], arr[endIndex]

	if start > median != (start > end){
		return startIndex
	} else if median > start != (median > end){
		return medianIndex
	} else {
		return endIndex
	}

}