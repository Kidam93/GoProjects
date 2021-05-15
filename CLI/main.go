package main

import (
	"fmt"
	"log"
	"os"
	// "time"
	"flag"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE IF NOT EXISTS entry
(
	id integer primary key auto_increment,
	title text,
	definition text,
	created_at date
)
`

type dbStore struct {
	db *sqlx.DB
}

type Entry struct {
	Id int
	Title string
	Definition string
	// CreatedAt time.Time
}

func main() {
	fmt.Println("HELLO")
	action := flag.String("action", "list", "Action to perform on the dictionnary")

	db := dbStore{}
	err := db.Open()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// go build -o cli
	// ./cli -action list
	flag.Parse()
	switch *action {
	case "list":
		actionList(&db)
	case "action":
		fmt.Println("case action.....")
	case "add": 
		actionAdd(&db, flag.Args())
	default:
		fmt.Printf("unknown action: %v", *action)
	}

	// debug
	// actionList(&db)
}

func (store *dbStore) Open() error {
	db, err := sqlx.Connect("mysql", "root:@/golangcli")
	if err != nil {
		return err
	}
	log.Println("Connected to DB")
	db.MustExec(schema)
	store.db = db
	return nil
}

func (store *dbStore) Close() error {
	return store.db.Close()
}

func actionList(store *dbStore) {
	entry, err := store.List()
	if err != nil {
		log.Fatal("error to list")
	}
	fmt.Printf("%v", entry)
}

func actionAdd(store *dbStore, args []string) {
	// call method add
	title := args[0]
	definition := args[1]
	store.Add(title, definition)
	fmt.Println("Added to entry !")
}

func actionDefine() {
	// call method get
}

func actionRemove() {
	// call method remove
}