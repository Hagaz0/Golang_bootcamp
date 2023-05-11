package main

import (
	"container/heap"
	"fmt"
	"os"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func (h PresentHeap) Len() int { return len(h) }
func (h PresentHeap) Less(i, j int) bool {
	if h[i].Value != h[j].Value {
		return h[i].Value > h[j].Value
	}
	return h[i].Size < h[j].Size
}

func (h PresentHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PresentHeap) Push(x interface{}) {
	item := x.(*Present)
	*h = append(*h, *item)
}

func (h *PresentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func getNCoolestPresents(h *PresentHeap, n int) {
	if n > h.Len() || n < 1 {
		fmt.Println("Недопустимый аргумент")
		os.Exit(1)
	}
	heap.Init(h)
	for i := 0; i < n; i++ {
		present := heap.Pop(h).(Present)
		fmt.Printf("{%d %d}\n", present.Value, present.Size)
	}
}

func main() {
	qwerty := &PresentHeap{
		{5, 1},
		{4, 5},
		{3, 1},
		{5, 2},
	}
	getNCoolestPresents(qwerty, 2)
}
