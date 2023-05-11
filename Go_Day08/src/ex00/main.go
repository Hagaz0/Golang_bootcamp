package main

import (
	"errors"
	"fmt"
)

func getElement(arr []int, idx int) (int, error) {
	if idx < 0 {
		return 0, errors.New(fmt.Sprintf("Индекс меньше нуля"))
	}
	if len(arr)-1 < idx {
		return 0, errors.New(fmt.Sprintf("Выход за пределы массива"))
	}
	if len(arr) == 0 {
		return 0, errors.New(fmt.Sprintf("Пустой массив"))
	}

	for i, e := range arr {
		if i == idx {
			return e, nil
		}
	}
	return 0, errors.New(fmt.Sprintf("Что-то пошло не так"))
}

func main() {
	var arr = []int{1, 2, 3}
	q, err := getElement(arr, 2)
	fmt.Println(q, err)
}
