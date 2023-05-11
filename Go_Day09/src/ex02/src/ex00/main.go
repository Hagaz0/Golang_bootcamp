package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepSort(arr []int) <-chan int {
	var wg sync.WaitGroup
	ch := make(chan int, len(arr))
	for _, v := range arr {
		wg.Add(1)
		go func(v int) {
			time.Sleep(time.Duration(v) * time.Second)
			ch <- v
			wg.Done()
		}(v)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

func main() {
	var arr = []int{2, 3, 1, 6, 4, 9, 7}
	ch := sleepSort(arr)
	for v := range ch {
		fmt.Print(v, " ")
	}
}
