package main

import (
	"fmt"
	"sync"
)

func multiplex(cs ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup

	out := make(chan interface{})

	send := func(c <-chan interface{}) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan any, 3)
	ch1 <- 1
	ch1 <- "2"
	ch1 <- 3.0
	close(ch1)

	ch2 := make(chan interface{}, 3)
	ch2 <- 1
	ch2 <- "4"
	ch2 <- 3.1
	close(ch2)
	res := multiplex(ch1, ch2)
	for {
		d, ok := <-res
		if !ok {
			break
		}
		fmt.Println(d)
	}
}
