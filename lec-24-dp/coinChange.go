package lec24dp

import "math"

func memo(coinFc func([]int, int)(int))(func([]int, int)(int)){
	cache := map[int]int{}

	memCoinFc := func(coins []int, amount int)(int){
		val, ok := cache[amount]; if ok {
			return val
		}

		res := coinFc(coins, amount)
		cache[amount] = res
		return res
	}

	return memCoinFc
}

func coinChange(coins []int, amount int) int {
	var memCoinFc func([]int, int) int

	memCoinFc = memo(func(coins []int, amount int) int {
		if amount < 0 {
			return math.MaxInt - 1 
		}

		if amount == 0 {
			return 0
		}

		
		minCoins := math.MaxInt - 1

		for _, coin := range coins {
			cnt := memCoinFc(coins, amount-coin)
			if cnt != math.MaxInt-1 {
				minCoins = min(minCoins, 1+cnt)
			}
		}

		return minCoins
	})

	res := memCoinFc(coins, amount)

	if res == math.MaxInt-1 {
		return -1
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}