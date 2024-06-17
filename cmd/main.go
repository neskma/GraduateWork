package main

import (
	"log"
	"my-go-web-scraper/config"
	"my-go-web-scraper/internal/downloader"
	"my-go-web-scraper/internal/parser"
	"my-go-web-scraper/internal/storage"
	"my-go-web-scraper/internal/utils"
	"my-go-web-scraper/internal/web"
	"my-go-web-scraper/models"
	"sync"
)

func main() {
	cfg := config.GetConfig()
	urls := []string{
		cfg.TargetSite.URL,
		// Добавьте другие URL для параллельного парсинга
	}

	var wg sync.WaitGroup
	dataChannel := make(chan []models.Book) // Используем models.Book

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			processURL(url, dataChannel)
		}(url)
	}

	go func() {
		wg.Wait()
		close(dataChannel)
	}()

	var allBooks []models.Book // Используем models.Book
	for books := range dataChannel {
		allBooks = append(allBooks, books...)
	}

	err := storage.SaveBooks(allBooks)
	utils.CheckErr(err, "Ошибка сохранения данных")

	log.Println("Парсинг завершен успешно!")

	// Запуск веб-сервера
	web.StartWebServer()
}

func processURL(url string, dataChannel chan<- []models.Book) { // Используем models.Book
	html, err := downloader.DownloadPage(url)
	if err != nil {
		log.Printf("Ошибка загрузки страницы %s: %v", url, err)
		return
	}

	books, err := parser.ParseBooks(html)
	if err != nil {
		log.Printf("Ошибка парсинга страницы %s: %v", url, err)
		return
	}

	dataChannel <- books
}
