package main

import (
	"fmt"
	"sort"
)

type Present struct {
	Value int
	Size int
}

func grapPresents(sl []Present, space int) []Present {
	sort.Slice(sl, func(i, j int) bool {
		return float64(sl[i].Value)/float64(sl[i].Size) > float64(sl[j].Value)/float64(sl[j].Size)
	})
	var res []Present
	for _, v := range sl {
		if v.Size <= space {
			res = append(res, v)
			space -= v.Size
		}
	}
	return res
}

func main() {
	sl := []Present{
		{10, 2},
		{20, 4},
		{15, 3},
		{5, 1},
	}
	fmt.Println(grapPresents(sl, 5))
}