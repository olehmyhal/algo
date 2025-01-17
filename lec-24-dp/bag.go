package lec24dp

import "fmt"

func bag(weightSl, priceSl []int, weight int) {
	n := len(priceSl) - 1

	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, weight+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= weight; j++ {
			if j >= weightSl[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weightSl[i]]+priceSl[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	tmp := weight
	ans := []int{}
	for i := n; i > 0; i-- {
		if dp[i][tmp] != dp[i-1][tmp] {
			ans = append(ans, weightSl[i])
			tmp -= weightSl[i]
		}
	}

	for _, row := range dp {
		fmt.Printf("%v \n", row)
	}

	fmt.Printf("ANSWER: %v", ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
