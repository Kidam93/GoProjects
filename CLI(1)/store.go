package main

import (
	"log"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	Open()
	Close()
}

type dbStore struct {
	db *sqlx.DB
}

func (store *dbStore) Open() error {
	db, err := sqlx.Connect("mysql", "root:@/golangcli")
	if err != nil {
		log.Println("Connected to DB")
	}
	store.db = db
	return nil
}

func (store *dbStore) Close() error {
	return store.db.Close()
}