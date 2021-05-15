package main

import (
	example "cli/example"
	// "fmt"
	// "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
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
	
	example.Web()
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