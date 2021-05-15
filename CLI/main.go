package main

import (
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type dbStore struct {
	db *sqlx.DB
}

func main() {
	fmt.Println("jjjjjjjjjjjjjjjd")
	db := dbStore{}
	err := db.Open()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func (store *dbStore) Open() error {
	db, err := sqlx.Connect("mysql", "root:@/golangcli")
	if err != nil {
		return err
	}
	log.Println("Connected to DB")
	store.db = db
	return nil
}

func (store *dbStore) Close() error {
	return store.db.Close()
}