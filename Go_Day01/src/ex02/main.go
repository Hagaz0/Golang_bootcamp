package main

import (
	"bufio"
	"fmt"
	"os"
)

func check_type_file(str string) {
	if str != ".txt" {
		fmt.Println("Файл отсутствует")
		os.Exit(1)
	}
}

func parse_args() map[string]string {
	if len(os.Args[1:]) != 4 {
		fmt.Println("Введите названия файлов с флагами --old или --new перед ними")
		os.Exit(1)
	}
	files := map[string]string{
		"old": "",
		"new": "",
	}
	for i, e := range os.Args[1:] {
		if i%2 == 0 && (e == "--old" || e == "--new") {
			if files[e[2:]] == "" {
				files[e[2:]] = os.Args[2+i]
			} else {
				fmt.Println("Не может быть двух new или old аргументов")
				os.Exit(1)
			}
		} else if i%2 == 1 {
			continue
		} else {
			fmt.Println("Введите названия файлов с флагами --old или --new перед ними")
			os.Exit(1)
		}
	}
	len_old := len(files["old"])
	len_new := len(files["new"])
	if len_old < 5 || len_new < 5 {
		fmt.Println("Введено неккоректное значение")
		os.Exit(1)
	}
	if files["old"] == files["new"] {
		fmt.Println("Введен один и тот же файл")
		os.Exit(1)
	}
	check_type_file(files["old"][len_old-4:])
	check_type_file(files["new"][len_new-4:])
	return files
}

func calculate() {
	files := parse_args()
	fold, err := os.Open(files["old"])
	if err != nil {
		panic(err)
	}
	defer fold.Close()

	fnew, err := os.Open(files["new"])
	if err != nil {
		panic(err)
	}
	defer fnew.Close()
	oldsc := bufio.NewScanner(fold)
	newsc := bufio.NewScanner(fnew)
	var old_lines, new_lines []string
	for oldsc.Scan() {
		old_lines = append(old_lines, oldsc.Text())
	}
	for newsc.Scan() {
		new_lines = append(new_lines, newsc.Text())
	}
	for _, enew := range new_lines {
		flag := 0
		for _, eold := range old_lines {
			if enew == eold {
				flag = 1
				break
			}
		}
		if flag == 0 {
			fmt.Printf("ADDED %s\n", enew)
		}
	}
	for _, eold := range old_lines {
		flag := 0
		for _, enew := range new_lines {
			if enew == eold {
				flag = 1
				break
			}
		}
		if flag == 0 {
			fmt.Printf("REMOVED %s\n", eold)
		}
	}
}

func main() {
	calculate()
}
