package main

import (
	"fmt"
	"log"
	"os"
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
	fmt.Println(".....................................")
	fmt.Println("WELCOME CLI")
	fmt.Println("show actions with: ./cli -action help")
	fmt.Println(".....................................")

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
	case "addmany":
		actionAddMany(&db, flag.Args())
	case "view":
		actionGet(&db, flag.Args())
	case "update":
		actionUpdate(&db, flag.Args())
	case "delete":
		actionRemove(&db, flag.Args())
	case "deleteall":
		actionRemoveAll(&db)
	case "help":
		actionHelp()
	default:
		fmt.Printf("unknown action: %v", *action)
	}
}

func (store *dbStore) Open() error {
	db, err := sqlx.Connect("mysql", "root:@/golangcli")
	if err != nil {
		return err
	}
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

func actionAddMany(store *dbStore, args []string) {
	title := args[0]
	definition := args[1]
	store.AddMany(title, definition)
	fmt.Println("Added more to entry !")
}

func actionGet(store *dbStore, args []string) {
	title := args[0]
	entry, err := store.Get(title)
	if err != nil {
		log.Fatal("error to view")
	}
	fmt.Printf("%v", entry)
}

func actionUpdate(store *dbStore, args []string) {
	title := args[0]
	definition := args[1]
	store.Update(title, definition)
	fmt.Println("Updated definition to entry !")
}

func actionRemove(store *dbStore, args []string) {
	title := args[0]
	store.Remove(title)
	fmt.Println("Removed to entry !")
}

func actionRemoveAll(store *dbStore) {
	store.RemoveAll()
	fmt.Println("Removed all to entry !")
}

func actionHelp() {
	fmt.Println("-------------------------------------")
	fmt.Println("./cli -action help")
	fmt.Println("./cli -action list")
	fmt.Println("./cli -action action")
	fmt.Println("./cli -action add title "+"definition")
	fmt.Println("./cli -action addmany title "+"definition")
	fmt.Println("./cli -action view title")
	fmt.Println("./cli -action update title "+"definition")
	fmt.Println("./cli -action delete title")
	fmt.Println("./cli -action deleteall")

	fmt.Println("-------------------------------------")
}