package main

import (
	"fmt"
)

type mas []int

func test(val1, val2 []int) bool {
	if len(val1) != len(val2) {
		return false
	}
	for i := 0; i < len(val1); i++ {
		if val1[i] != val2[i] {
			return false
		}
	}
	return true
}

func main() {
	val := []int{1, 2, 6, 9, 10, 50, 32, 45, 0, 87}
	coins := []mas{
		{1, 2, 3},
		{5, 6, 2, 1},
		{4, 5, 3, 7},
		{10, 100, 5, 4, 2, 3},
		{5, 4, 7, 4, 2},
		{1, 2, 9, 6},
		{5, 200, 1000},
		{10000},
		{9},
		{1, 2, 100, 80, 4},
	}
	count := 0
	for i, _ := range val {
		count++
		res1, res2 := minCoins(val[i], coins[i]), minCoins2(val[i], coins[i])
		if test(res1, res2) {
			fmt.Printf("%d test: Success\n", count)
		} else {
			fmt.Printf("%d test: Failed\nNormal output: ", count)
			fmt.Println(res2)
			fmt.Print("Your output: ")
			fmt.Println(res1)
		}
	}
}
