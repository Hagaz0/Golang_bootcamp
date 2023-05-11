package main

import "sort"

func minCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

func minCoins2(val int, coins []int) []int {
	res := make([]int, 0)
	sort.Ints(coins)
	for i := len(coins) - 1; i >= 0; i-- {
		if val == 0 {
			break
		}
		if val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
	}
	return res
}
