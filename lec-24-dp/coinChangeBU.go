package lec24dp

import "math"

func coinChangeBU(coins []int, amount int) int {
    cache := make([]int, amount + 1)
    sentinel := math.MaxInt - 1

	cache[0] = 0

    for currAmount := 1; currAmount < len(cache); currAmount++ {
        for _, coin := range coins {
            if currAmount >= coin {
                cache[currAmount] = min(cache[currAmount], cache[currAmount-coin]+1)
            }
        }
    }

    if cache[amount] == sentinel {
        return -1
    }

    return cache[amount]
}

func minBU(a, b int) int {
	if a < b {
		return a
	}
	return b
}