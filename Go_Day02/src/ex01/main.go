package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func parse_args() (map[string]bool, []string) {
	flags := map[string]bool{
		"w": false,
		"l": false,
		"m": false,
	}
	var files []string
	for _, e := range os.Args[1:] {
		if e == "-w" {
			flags["w"] = true
		} else if e == "-l" {
			flags["l"] = true
		} else if e == "-m" {
			flags["m"] = true
		} else {
			files = append(files, e)
		}
	}
	if !flags["w"] && !flags["l"] && !flags["m"] {
		flags["w"] = true
	} else if (flags["w"] && flags["l"]) || (flags["w"] && flags["m"]) || (flags["l"] && flags["m"]) {
		fmt.Println("Введено более одного флага")
		os.Exit(1)
	}
	if len(files) == 0 {
		fmt.Println("Введите файл")
		os.Exit(1)
	}
	for _, e := range files {
		if _, err := os.Stat(e); os.IsNotExist(err) {
			fmt.Println("Файл", e, "не существует")
			os.Exit(1)
		}
	}
	return flags, files
}

func parse_words(files []string) {
	wg := new(sync.WaitGroup)
	for _, e := range files {
		wg.Add(1)
		go func(e string) {
			file, err := os.Open(e)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanWords)
			var count int
			for scanner.Scan() {
				count++
			}
			fmt.Printf("%d\t%s\n", count, e)
			wg.Done()
		}(e)
	}
	wg.Wait()
}

func parse_lines(files []string) {
	wg := new(sync.WaitGroup)
	for _, e := range files {
		wg.Add(1)
		go func(e string) {
			file, err := os.Open(e)
			if err != nil {
				panic(err)
			}
			var count int
			defer file.Close()
			sc := bufio.NewScanner(file)
			for sc.Scan() {
				count++
			}
			fmt.Printf("%d\t%s\n", count, e)
			wg.Done()
		}(e)
	}
	wg.Wait()
}

func parse_symbols(files []string) {
	wg := new(sync.WaitGroup)
	for _, e := range files {
		wg.Add(1)
		go func(e string) {
			data, err := ioutil.ReadFile(e)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%d\t%s\n", len(data), e)
			wg.Done()
		}(e)
	}
	wg.Wait()
}

func main() {
	flags, files := parse_args()
	if flags["w"] {
		parse_words(files)
	} else if flags["l"] {
		parse_lines(files)
	} else if flags["m"] {
		parse_symbols(files)
	}
}
