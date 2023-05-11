package main

import (
	"fmt"
	"reflect"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(input interface{}) {
	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			value := v.Field(i)
			fmt.Printf("%v: %v\n", field.Name, value.Interface())
		}
	} else {
		fmt.Println("Input is not a struct")
	}
}

func main() {
	up := UnknownPlant{"lepestok", "pelmen", 123}
	aup := AnotherUnknownPlant{15, "aboba", 4567}
	describePlant(up)
	fmt.Println()
	describePlant(aup)
}
