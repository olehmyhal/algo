package tasks

import (
	"fmt"
	"math"
)

//1tc failing
func bruh4(){
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	points := make([][2]int, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)

		points[i] = [2]int{x, y}
	}

	sortedByCoordinates := sort2(points)

	sortedPoints := sort(sortedByCoordinates)

	for _, coord := range sortedPoints[:k] {
		fmt.Printf("%v %v\n", coord[0], coord[1])
	}

}

func sort(arr [][2]int)[][2]int{
	merge_sort(0, len(arr), arr)

	return arr
}

func merge_sort(start, end int, arr [][2]int){
	if end - start < 2 {
		return
	}

	m := (start + end) / 2

	merge_sort(start, m, arr)
	merge_sort(m, end, arr)
	merge(start, end, m, arr)
}

func merge(start, end, m int, arr [][2]int){
	aux := [][2]int{}

	for i := start; i < m; i++ {
		aux = append(aux, arr[i])
	}

	i, j, k := 0, m, start

	for i < len(aux) && j < end {
		if calcLength(aux[i]) < calcLength(arr[j]) {
			arr[k] = aux[i]
			i++
		} else if calcLength(aux[i]) > calcLength(arr[j]) {
			arr[k] = arr[j]
			j++
		} else {
			if aux[i][0] < arr[j][0] {
				arr[k] = aux[i]
				i++
			} else if aux[i][0] > arr[j][0] {
				arr[k] = arr[j]
				j++
			} 
		}
		k++
	}

	for i < len(aux) {
		arr[k] = aux[i]
		i++
		k++
	}

	for j < end {
		arr[k] = arr[j]
		j++
		k++
	}
}

func calcLength(point [2]int)float64{
	x, y := float64(point[0]), float64(point[1])
	return float64(math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2)))
}

func sort2(arr [][2]int)[][2]int{
	merge_sort2(0, len(arr), arr)

	return arr
}

func merge_sort2(start, end int, arr [][2]int){
	if end - start < 2 {
		return
	}

	m := (start + end) / 2

	merge_sort2(start, m, arr)
	merge_sort2(m, end, arr)
	merge2(start, end, m, arr)
}

func merge2(start, end, m int, arr [][2]int){
	aux := [][2]int{}

	for i := start; i < m; i++ {
		aux = append(aux, arr[i])
	}

	i, j, k := 0, m, start

	for i < len(aux) && j < end {
		if aux[i][0] < arr[j][0] {
			arr[k] = aux[i]
			i++
		} else if aux[i][0] > arr[j][0] {
			arr[k] = arr[j]
			j++
		} else {
			if aux[i][1] < arr[j][1] {
				arr[k] = aux[i]
				i++
			} else {
				arr[k] = arr[j]
				j++
			}
		}
		k++
	}

	for i < len(aux) {
		arr[k] = aux[i]
		i++
		k++
	}

	for j < end {
		arr[k] = arr[j]
		j++
		k++
	}
}
