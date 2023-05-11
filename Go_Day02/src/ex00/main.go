package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func parse_args() {
	if len(os.Args[1:]) > 6 {
		fmt.Printf("Введите до 6 аргументов\n")
		os.Exit(1)
	}
	flags := map[string]bool{
		"sl":  false,
		"d":   false,
		"f":   false,
		"ext": false,
	}
	datas := map[string]string{
		"path": "",
		"type": "",
	}
	for _, e := range os.Args[1:] {
		if e == "-sl" {
			flags["sl"] = true
		} else if e == "-d" {
			flags["d"] = true
		} else if e == "-f" {
			flags["f"] = true
		} else if e == "-ext" {
			flags["ext"] = true
		} else if string(e[0]) == "/" {
			datas["path"] = e
		} else if string(e[0]) == "'" && string(e[len(e)-1]) == "'" {
			datas["type"] = "." + e[1:len(e)-1]
		} else {
			fmt.Printf("Введен некорректный аргумент %s\n", e)
			os.Exit(1)
		}
	}
	if flags["ext"] && (!flags["f"] || flags["sl"] || flags["d"]) {
		fmt.Printf("Невозможна работа флага -ext при выключенном флаге -f или включенных других флагах\n")
		os.Exit(1)
	}
	if flags["ext"] && datas["type"] == "" {
		fmt.Printf("Отсутствует шаблон поиска по типу файла\n")
		os.Exit(1)
	}
	if datas["path"] == "" {
		fmt.Printf("Введите путь\n")
		os.Exit(1)
	}
	if datas["type"] != "" && flags["ext"] {
		fmt.Printf("Введите -ext для считывания шаблона\n")
		os.Exit(1)
	}
	if !flags["f"] && !flags["d"] && !flags["sl"] {
		flags["f"] = true
		flags["d"] = true
		flags["sl"] = true
	}
	listDirByWalk(flags, datas)
}

func listDirByWalk(flags map[string]bool, datas map[string]string) {
	filepath.Walk(datas["path"], func(wPath string, info os.FileInfo, err error) error {

		// Обход директории без вывода
		if wPath == datas["path"] {
			return nil
		}

		fileinfo, err := os.Lstat(wPath)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("%s -> [broken]\n", wPath)
			} else {
				fmt.Println(err)
			}
		} else if fileinfo.Mode().IsRegular() && flags["f"] {
			if flags["ext"] {
				if filepath.Ext(wPath) == datas["type"] {
					fmt.Println(wPath)
				}
			} else {
				//fmt.Println("Это обычный файл")
				fmt.Println(wPath)
			}
		} else if fileinfo.Mode()&os.ModeSymlink != 0 && flags["sl"] {
			linkPath, err := os.Readlink(wPath)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println("Это символическая ссылка")
			fmt.Printf("%s -> %s\n", wPath, linkPath)
		} else if fileinfo.Mode().IsDir() && flags["d"] {
			fmt.Println(wPath)
		}
		return nil
	})
}

func main() {
	parse_args()
}
