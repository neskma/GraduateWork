package tests

import (
	"my-go-web-scraper/internal/downloader"
	"testing"
)

func TestDownloadPage(t *testing.T) {
	url := "https://books.toscrape.com/"
	_, err := downloader.DownloadPage(url)
	if err != nil {
		t.Fatalf("Ошибка загрузки страницы: %v", err)
	}
}
