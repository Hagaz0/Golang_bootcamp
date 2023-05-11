package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func sleepSort(arr []int) chan int {
	for _, e := range arr {
		if e < 0 {
			fmt.Println("Недопустимое значение в срезе")
			os.Exit(1)
		}
	}
	var wg sync.WaitGroup
	ch := make(chan int)
	for _, e := range arr {
		wg.Add(1)
		go func(e int) {
			defer wg.Done()
			time.Sleep(time.Duration(e) * time.Second)
			ch <- e
		}(e)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

func main() {
	var arr = []int{3, 2, 4, 6, 5, 0, 1}
	chanel := sleepSort(arr)
	for e := range chanel {
		fmt.Println(e)
	}
}
