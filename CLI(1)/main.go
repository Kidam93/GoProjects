package main

import (
	// dictionnary "cli/dictionnary"
	// example "cli/example"
	// "fmt"
	// "os"
	// "fmt"
	// "database/sql"
	"os"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("CLI")

	srv := newServer()
	srv.store = &dbStore{}
	err := srv.store.Open()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	defer srv.store.Close()

	// srv := newServer()
	// srv.store = &dbStore{}
	

	// fmt.Println("database..")

	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golangcli")
	// if err != nil {
    //     panic(err.Error())
    // }
	// defer db.Close()

	// insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
    // if err != nil {
    //     panic(err.Error())
    // }
    // defer insert.Close()
	
	// example.Web()
}

func actionList() {
	// call method list
}

func actionAdd() {
	// call method add
}

func actionDefine() {
	// call method get
}

func actionRemove() {
	// call method remove
}