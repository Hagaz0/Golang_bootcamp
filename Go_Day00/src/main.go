package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

//  Ввод массива

func Input_mas() []int {
	var mas []int
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Bytes()
		if len(txt) != 0 {
			if i, err := strconv.Atoi(string(txt)); err != nil || i < -100000 || i > 100000 {
				fmt.Println("Введите корректные значения в диапазоне от -100000 до 100000")
			} else {
				mas = append(mas, i)
			}
		} else if len(txt) == 0 {
			break
		}
	}
	sort.Ints(mas)
	return mas
}

func Parse_args(Args *map[string]bool) {
	q := len(os.Args[1:])
	if q != 0 && q < 5 {
		for _, e := range os.Args[1:] {
			switch e {
			case "mean":
				(*Args)["mean"] = true
			case "median":
				(*Args)["median"] = true
			case "mode":
				(*Args)["mode"] = true
			case "sd":
				(*Args)["sd"] = true
			default:
				fmt.Println("Введен некорректный аргумент. Корректные аргументы: mean, median, mode, sd")
				os.Exit(1)
			}
		}
	} else if q > 4 {
		fmt.Println("Введено неккоректное число аргументов. Введите до 4х аргументов")
		os.Exit(1)
	} else {
		(*Args)["mean"] = true
		(*Args)["median"] = true
		(*Args)["mode"] = true
		(*Args)["sd"] = true
	}
}

func mean(mas []int) float64 {
	var res float64
	count := 0
	for _, e := range mas {
		res += float64(e)
		count++
	}
	return res / float64(count)
}

func median(mas []int) {
	if len(mas)%2 == 1 {
		fmt.Printf("Median: %d\n", mas[len(mas)/2])
	} else {
		fmt.Printf("Median: %.02f\n", (float64(mas[len(mas)/2])+float64(mas[len(mas)/2-1]))/2.0)
	}
}

func mode(mas []int) {
	numbers := make(map[int]int)
	max := 0
	for _, e := range mas {
		numbers[e]++
	}
	for _, e := range numbers {
		if e > max {
			max = e
		}
	}
	for i, e := range numbers {
		if max == e {
			fmt.Printf("Mode: %d\n", i)
			break
		}
	}
}

func sd(mas []int) {
	if len(mas) == 1 {
		fmt.Printf("SD: NaN\n")
	} else {
		mean := mean(mas)
		var q, res float64
		for _, e := range mas {
			q = (float64(e) - mean) * (float64(e) - mean)
			res += q
		}
		res = math.Sqrt(res / float64(len(mas)-1))
		fmt.Printf("SD: %.02f\n", res)
	}
}

func calculate(mas []int, Args map[string]bool) {
	for i, e := range Args {
		if e {
			switch i {
			case "mean":
				fmt.Printf("Mean: %.02f\n", mean(mas))
			case "median":
				median(mas)
			case "mode":
				mode(mas)
			case "sd":
				sd(mas)
			}
		}
	}
}

func main() {
	Args := map[string]bool{
		"mean":   false,
		"median": false,
		"mode":   false,
		"sd":     false,
	}
	Parse_args(&Args)
	Sequence := Input_mas()
	if len(Sequence) != 0 {
		calculate(Sequence, Args)
	} else {
		fmt.Println("Пустая последовательность")
	}
}
