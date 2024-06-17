package parser

import (
	"my-go-web-scraper/models"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseBooks(html string) ([]models.Book, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var books []models.Book
	doc.Find(".product_pod").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3 a").AttrOr("title", "")
		price := s.Find(".price_color").Text()
		rating := s.Find(".star-rating").AttrOr("class", "")
		rating = strings.Replace(rating, "star-rating ", "", 1)

		books = append(books, models.Book{Title: title, Price: price, Rating: rating})
	})

	return books, nil
}
