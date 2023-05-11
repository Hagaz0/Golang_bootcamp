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
		Ingredients struct {
			Item []struct {
				Itemname  string `xml:"itemname"`
				Itemcount string `xml:"itemcount"`
				Itemunit  string `xml:"itemunit"`
			} `xml:"item"`
		} `xml:"ingredients"`
	} `xml:"cake"`
}

type data struct {
	xml_file  xml_cake
	json_file recipes
	type_old  string
}

func parse_xml(filename string) xml_cake {
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
	return xml_result
}

func parse_json(filename string) recipes {
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
	return json_result
}

func check_type_file(str string) {
	switch str {
	case ".xml":
	case "json":
	default:
		fmt.Println("Файл отсутствует")
		os.Exit(1)
	}
}

func parse_args() (map[string]string, string) {
	count_args := len(os.Args[1:])
	if count_args != 4 {
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
	type_old := files["old"][len_old-4:]
	type_new := files["new"][len_new-4:]
	check_type_file(type_old)
	check_type_file(type_new)
	if type_old == type_new {
		fmt.Println("Введите файлы с разными типами (xml и json)")
		os.Exit(1)
	}
	return files, type_old
}

func calculate(dataf data) {
	for _, enew := range dataf.json_file.Cake {
		flag := 0
		for _, eold := range dataf.xml_file.Cake {
			if eold.Name == enew.Name {
				flag = 1
				break
			}
		}
		if flag == 0 {
			fmt.Printf("ADDED cake \"%s\"\n", enew.Name)
		}
	}
	for _, eold := range dataf.xml_file.Cake {
		flag := 0
		for _, enew := range dataf.json_file.Cake {
			if eold.Name == enew.Name {
				flag = 1
				break
			}
		}
		if flag == 0 {
			fmt.Printf("REMOVED cake \"%s\"\n", eold.Name)
		}
	}
	for _, eold := range dataf.xml_file.Cake {
		for _, enew := range dataf.json_file.Cake {
			if eold.Name == enew.Name && eold.Stovetime != enew.Time {
				fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", eold.Name, enew.Time, eold.Stovetime)
			}
		}
	}
	for _, eold := range dataf.xml_file.Cake {
		for _, enew := range dataf.json_file.Cake {
			if eold.Name == enew.Name {
				for _, ing_new := range enew.Ingredients {
					flag := 0
					for _, ing_old := range eold.Ingredients.Item {
						if ing_old.Itemname == ing_new.Ingredient_name {
							flag = 1
							break
						}
					}
					if flag == 0 {
						fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", ing_new.Ingredient_name, eold.Name)
					}
				}

				for _, ing_old := range eold.Ingredients.Item {
					flag := 0
					for _, ing_new := range enew.Ingredients {
						if ing_old.Itemname == ing_new.Ingredient_name {
							flag = 1
							break
						}
					}
					if flag == 0 {
						fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", ing_old.Itemname, eold.Name)
					}
				}
				for _, ing_new := range enew.Ingredients {
					for _, ing_old := range eold.Ingredients.Item {
						if ing_old.Itemname == ing_new.Ingredient_name && ing_old.Itemcount != ing_new.Ingredient_count {
							fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake"+
								"  \"%s\" - \"%s\" instead of"+
								" \"%s\"\n", ing_new.Ingredient_name, eold.Name, ing_new.Ingredient_count, ing_old.Itemcount)
						}
					}
				}
				for _, ing_new := range enew.Ingredients {
					for _, ing_old := range eold.Ingredients.Item {
						if ing_old.Itemname == ing_new.Ingredient_name && ing_old.Itemunit != ing_new.Ingredient_unit {
							if ing_new.Ingredient_unit == "" {
								fmt.Printf("REMOVED unit \"%s\" for ingredient"+
									" \"%s\" for cake  \"%s\"\n", ing_old.Itemunit, ing_old.Itemname, eold.Name)
							} else if ing_old.Itemunit == "" {
								fmt.Printf("ADDED unit \"%s\" for ingredient"+
									" \"%s\" for cake  \"%s\"\n", ing_new.Ingredient_unit, ing_old.Itemname, eold.Name)
							} else {
								fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  "+
									"\"%s\" - \"%s\" instead of "+
									"\"%s\"\n", ing_new.Ingredient_name, eold.Name, ing_new.Ingredient_unit, ing_old.Itemunit)
							}
						}
					}
				}
			}
		}
	}
}

func main() {
	files, type_old := parse_args()
	var data_files data
	data_files.type_old = type_old
	if type_old == ".xml" {
		data_files.xml_file = parse_xml(files["old"])
		data_files.json_file = parse_json(files["new"])
		calculate(data_files)
	} else {
		fmt.Println("Old база данных должна быть формата xml")
		os.Exit(1)
	}
}
