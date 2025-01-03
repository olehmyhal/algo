package tasks

import "fmt"

func bruh3(){
    var m, n, target int
    fmt.Scanf("%d %d %d", &m, &n, &target)

    matrix := make([][]int, m)
    for i := 0; i < m; i++ {
        matrix[i] = make([]int, n)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            fmt.Scan(&matrix[i][j])
        }
    }

    firstValues := make([]int, m)

    for i, row := range matrix {
        firstValues[i] = row[0]
    }

    i, j := -1, m

    for j - i > 1 {
        med := (j + i) / 2

        if firstValues[med] > target {
            j = med
        } else {
            i = med
        }
    }

	if i < 0 {
		fmt.Printf("%v", 0)
		return
	}

    choosenRow := matrix[i]

    k, l := -1, n

    for l - k > 1 {
        med := (k + l) / 2

        if choosenRow[med] > target {
            l = med
        } else {
            k = med
        }
    }

	if k < 0 {
		fmt.Printf("%v", 0)
		return
	}

    if choosenRow[k] == target {
        fmt.Printf("%v", 1)
    } else {
        fmt.Printf("%v", 0)
    }
}