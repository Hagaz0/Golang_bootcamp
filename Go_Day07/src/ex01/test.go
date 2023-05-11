package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type mas []int

func main() {
	myspeed := make([]int, 0)
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
	file, err := os.OpenFile("top10.txt", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		file, err = os.Create("top10.txt")
		if err != nil {
			log.Fatal(err)
		}
	}
	defer file.Close()
	count := 0
	for i, _ := range val {
		count++
		mystart := time.Now()
		minCoins2(val[i], coins[i])
		myend := time.Since(mystart)
		start := time.Now()
		minCoins(val[i], coins[i])
		end := time.Since(start)

		myspeed = append(myspeed, int(myend.Nanoseconds()))

		fmt.Printf("%d Test - My duraion: %s, Your duration: %s\n", count, myend.String(), end.String())
	}
	sort.Ints(myspeed)
	count = 0
	for i := len(myspeed) - 1; i >= 0; i-- {
		count++
		writer := bufio.NewWriter(file)
		str := fmt.Sprintf("%d - %d\n", count, myspeed[i])
		if err != nil {
			log.Fatal(err)
		}
		_, err := writer.WriteString(str)
		if err != nil {
			log.Fatal(err)
		}
		writer.Flush()
	}
}
