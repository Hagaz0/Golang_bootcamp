package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Ingredients struct {
	Ingredient_name  string
	Ingredient_count string
	Ingredient_unit  string
}

type Cake struct {
	Name        string
	Time        string
	Ingredients []Ingredients
}

// json struct
type recipes struct {
	Cake []Cake
}

// xml struct
type xml_cake struct {
	Cake []struct {
		Name        string `xml:"name"`
		Stovetime   string `xml:"stovetime"`
		Ingredients []struct {
			Item []struct {
				Itemname  string `xml:"itemname"`
				Itemcount string `xml:"itemcount"`
				Itemunit  string `xml:"itemunit"`
			} `xml:"item"`
		} `xml:"ingredients"`
	} `xml:"cake"`
}

func print_xml(filename string) {
	f_xml, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f_xml.Close()
	xml_data, err := ioutil.ReadAll(f_xml)
	if err != nil {
		log.Fatal(err)
	}
	var xml_result xml_cake
	xmlErr := xml.Unmarshal(xml_data, &xml_result)
	if xmlErr != nil {
		log.Fatal(xmlErr)
	}
	jsres, _ := json.MarshalIndent(xml_result, "", "    ")
	fmt.Printf("%s", jsres)
}

func print_json(filename string) {
	f_json, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f_json.Close()
	json_data, err := ioutil.ReadAll(f_json)
	if err != nil {
		log.Fatal(err)
	}
	var json_result recipes
	jsonErr := json.Unmarshal(json_data, &json_result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	xmlres, _ := xml.MarshalIndent(json_result, "", "    ")
	fmt.Printf("%s", xmlres)
}

func parse_args() {
	count_args := len(os.Args[1:])
	if count_args == 0 {
		fmt.Println("Аргумент отсутствует. Введите *.xml/*.json файл")
		os.Exit(1)
	} else if count_args > 1 {
		fmt.Println("Введите только один аргумент - название файла")
		os.Exit(1)
	} else {
		len_arg := len(os.Args[1])
		if len_arg < 5 {
			fmt.Println("Введено неккоректное значение")
			os.Exit(1)
		}
		type_file := os.Args[1][len(os.Args[1])-4:]
		switch type_file {
		case ".xml":
			print_xml(os.Args[1])
		case "json":
			print_json(os.Args[1])
		default:
			fmt.Println("Файл отсутствует")
			os.Exit(1)
		}
		fmt.Println(type_file)
	}
}

func main() {
	parse_args()
}
