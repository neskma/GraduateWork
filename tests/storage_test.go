package tests

import (
	"encoding/json"
	"my-go-web-scraper/internal/storage"
	"my-go-web-scraper/models"
	"os"
	"testing"
)

func TestSaveBooks(t *testing.T) {
	books := []models.Book{
		{Title: "Book 1", Price: "£23.88", Rating: "Three"},
		{Title: "Book 2", Price: "£19.99", Rating: "Four"},
	}

	err := storage.SaveBooks(books)
	if err != nil {
		t.Fatalf("Ошибка сохранения книг: %v", err)
	}

	file, err := os.Open("books.json")
	if err != nil {
		t.Fatalf("Ошибка открытия файла books.json: %v", err)
	}
	defer file.Close()

	var savedBooks []models.Book
	err = json.NewDecoder(file).Decode(&savedBooks)
	if err != nil {
		t.Fatalf("Ошибка декодирования JSON из файла: %v", err)
	}

	if len(savedBooks) != len(books) {
		t.Fatalf("Ожидаемое количество книг: %d, полученное: %d", len(books), len(savedBooks))
	}

	for i, book := range savedBooks {
		if book.Title != books[i].Title || book.Price != books[i].Price || book.Rating != books[i].Rating {
			t.Errorf("Ожидалось: %+v, получено: %+v", books[i], book)
		}
	}

	// Удаление файла после теста
	err = os.Remove("books.json")
	if err != nil {
		t.Fatalf("Ошибка удаления файла books.json: %v", err)
	}
}
