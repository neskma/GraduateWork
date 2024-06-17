package storage

import (
	"database/sql"
	"my-go-web-scraper/models"

	_ "github.com/mattn/go-sqlite3"
)

func SaveBooks(books []models.Book) error { // Используем models.Book
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		return err
	}
	defer db.Close()

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        price TEXT,
        rating TEXT
    );`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	insertBookQuery := `INSERT INTO books (title, price, rating) VALUES (?, ?, ?);`
	stmt, err := db.Prepare(insertBookQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, book := range books {
		_, err = stmt.Exec(book.Title, book.Price, book.Rating)
		if err != nil {
			return err
		}
	}

	return nil
}
