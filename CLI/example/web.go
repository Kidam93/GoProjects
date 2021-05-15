package example

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

var schema = `
CREATE TABLE IF NOT EXISTS book
(
	id integer primary key auto_increment,
	title text,
	author text,
	page_count integer
)
`

type Book struct{
	Id int
	Title string
	Author string
	PageCount int
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello gophers!")
}

func movies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "affichages des films...")
}

// http:localhost:9090/search?t=go&p=1
func search(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("t")
	p := r.URL.Query().Get("p")
	fmt.Fprintf(w, "Searching for term=%v. Current page=%v", t, p)
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "example/login.html")
		return
	case "POST":
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "ParseForm() failed. err=%v", err)
			return
		}
		fmt.Fprintf(w, "Go login POST. value=%v\n", r.PostForm)
		username := r.FormValue("username")
		password := r.FormValue("password")
		
		if username == "Go" && password == "password" {
			fmt.Fprintf(w, "OK !")
		} else {
			fmt.Fprintf(w, "Pas OK !")
		}
	}
}

func database() {
	db, err := sql.Open("mysql", "root:@/golangcli")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(schema)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()

	var id int
	var title string
	var author string
	var pageCount int

	rows, _ := db.Query("SELECT * FROM book")
	for rows.Next() {
		rows.Scan(&id, &title, &author, &pageCount)
		fmt.Printf("id=%v, title=%v, author=%v, pageCount=%v\n", id, title, author, pageCount)
	}

	rows, _ = db.Query("SELECT * FROM book")
	b := Book{}
	for rows.Next() {
		rows.Scan(&b.Id, &b.Title, &b.Author, &b.PageCount)
		fmt.Printf("book=%v\n", b)
	}

	fmt.Println("Ping to database successfull")

	b = Book{
		Title:		"Harry potter à l'école des soriciers",
		Author:		"J.K Rowling",
		PageCount: 308,
	}

	stmt, _ = db.Prepare("INSERT INTO book(title, author, page_count) VALUES (?, ?, ?)")
	_, err = stmt.Exec(b.Title, b.Author, b.PageCount)
	if err != nil {
		log.Fatal(err)
	}
}

func Web() {
	// http.HandleFunc("/", hello)
	// http.HandleFunc("/movies", movies)
	// http.HandleFunc("/search", search)
	// http.HandleFunc("/login", login)

	// if err := http.ListenAndServe(":9090", nil);
	// err != nil {
	// 	log.Fatal(err)
	// }

	database()
}