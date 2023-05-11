package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func parse_args() (string, []string) {
	var output_dir string
	var files []string
	for i, e := range os.Args[1:] {
		if e == "-a" {
			if i+1 == len(os.Args[1:]) {
				fmt.Println("Неккоректный ввод")
				os.Exit(1)
			}
			output_dir = os.Args[i+2]
			fileinfo, err := os.Lstat(output_dir)
			if err != nil {
				log.Fatal(err)
			}
			if !fileinfo.Mode().IsDir() {
				fmt.Printf("%s не является директорией", output_dir)
				os.Exit(1)
			}
		} else if filepath.Ext(e) == ".log" {
			files = append(files, e)
		}
	}
	if output_dir == "" && len(files) == 1 {
		output_dir = filepath.Dir(files[0])
	} else if output_dir == "" && len(files) > 1 {
		fmt.Println("При неуказанном каталоге вывода, введите 1 файл для архивирования")
		os.Exit(1)
	}
	return output_dir, files
}

func archive() {
	output_dir, files := parse_args()
	for _, e := range files {
		q := filepath.Base(e)
		output_file := q[:len(q)-4] + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ".tar"
		file, err := os.Open(e)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			log.Fatal(err)
		}

		outputFilePath := filepath.Join(output_dir, output_file)
		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer outputFile.Close()

		tarWriter := tar.NewWriter(outputFile)
		defer tarWriter.Close()

		header := &tar.Header{
			Name: fileInfo.Name(),
			Mode: int64(fileInfo.Mode()),
			Size: fileInfo.Size(),
		}
		err = tarWriter.WriteHeader(header)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(tarWriter, file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Файл успешно запакован в", outputFilePath)
	}
}

func main() {
	archive()
}
