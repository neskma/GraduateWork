package web

import (
	"database/sql"
	"html/template"
	"log"
	"my-go-web-scraper/models"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func StartWebServer() {
	http.HandleFunc("/", homeHandler)
	log.Println("Запуск веб-сервера на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT title, price, rating FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.Title, &book.Price, &book.Rating)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	tmpl, err := template.New("home").Parse(`
    <html>
        <head><title>Books</title></head>
        <body>
            <h1>Books</h1>
            <ul>
            {{range .}}
                <li>{{.Title}} - {{.Price}} - {{.Rating}}</li>
            {{end}}
            </ul>
        </body>
    </html>
    `)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, books)
}
