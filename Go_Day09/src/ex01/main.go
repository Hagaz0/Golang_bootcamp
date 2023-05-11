package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func crawlWeb(ctx context.Context, urls <-chan string) <-chan *string {
	var wg sync.WaitGroup
	results := make(chan *string)

	// Создаем пул горутин для обработки URL-адресов
	sem := make(chan struct{}, 8)

	for url := range urls {
		select {
		case <-ctx.Done():
			return results
		default:
			// Запускаем горутину для обработки URL-адреса
			sem <- struct{}{}
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				defer func() { <-sem }()
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("Error fetching %s: %s\n", url, err)
					return
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error reading body of %s: %s\n", url, err)
					return
				}
				result := string(body)
				results <- &result
			}(url)
		}
	}

	// Ожидаем завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}

func main() {
	// Создаем контекст для отмены операции
	ctx := context.Background()

	// Создаем каналы для входных URL-адресов и результатов
	urls := make(chan string)
	//results := make(chan *string)

	// Запускаем функцию crawlWeb в отдельной горутине
	results := crawlWeb(ctx, urls)

	// Отправляем несколько URL-адресов в канал urls
	//urls <- "https://www.example.com"
	urls <- "https://www.google.com"
	urls <- "https://www.yahoo.com"

	// Закрываем канал urls, чтобы crawlWeb завершил работу
	close(urls)

	// Читаем результаты из канала results и выводим их на экран
	for result := range results {
		if result != nil {
			fmt.Println(*result)
		} else {
			fmt.Println("Error fetching URL")
		}
	}
}
