package tests

import (
	"my-go-web-scraper/internal/parser"
	"my-go-web-scraper/models"
	"testing"
)

const sampleHTML = `
<html>
    <body>
        <div class="product_pod">
            <h3><a title="Book 1"></a></h3>
            <p class="price_color">£23.88</p>
            <p class="star-rating Three"></p>
        </div>
        <div class="product_pod">
            <h3><a title="Book 2"></a></h3>
            <p class="price_color">£19.99</p>
            <p class="star-rating Four"></p>
        </div>
    </body>
</html>
`

func TestParseBooks(t *testing.T) {
	books, err := parser.ParseBooks(sampleHTML)
	if err != nil {
		t.Fatalf("Ошибка парсинга HTML: %v", err)
	}

	expectedBooks := []models.Book{
		{Title: "Book 1", Price: "£23.88", Rating: "star-rating Three"},
		{Title: "Book 2", Price: "£19.99", Rating: "star-rating Four"},
	}

	if len(books) != len(expectedBooks) {
		t.Fatalf("Ожидаемое количество книг: %d, полученное: %d", len(expectedBooks), len(books))
	}

	for i, book := range books {
		if book.Title != expectedBooks[i].Title || book.Price != expectedBooks[i].Price || book.Rating != expectedBooks[i].Rating {
			t.Errorf("Ожидалось: %+v, получено: %+v", expectedBooks[i], book)
		}
	}
}
