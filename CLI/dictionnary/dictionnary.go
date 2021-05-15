package dictionnary

import (
	// "log"
	// "database/sql"
	// _ "github.com/lib/pq"
    // "github.com/jmoiron/sqlx"
	// _ "github.com/go-sql-driver/mysql"
)

type Store interface {
	Open()
	Close()
}

type Entry struct {
	
}

func (e Entry) Open() {

}

func (e Entry) Close() {

}