package main

import (
	"encoding/gob"
	"strings"
	// "flag"
	"fmt"
	"log"
	"time"
	"bytes"
	"github.com/boltdb/bolt"
)

type DB struct{
	db *bolt.DB
}

type Entry struct{
	Word string
	Definition string
	CreatedAt time.Time
}

func main(){
	// action := flag.String("action", "list", "Action to be perform on the dictionary")
	d, err := New() /* recuperation de la structure ( d := DB{ ...initialized value... } ) */
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("Struct DB: %v", d)
	d.CreateBuckets()

	/* without CLI */
	d.Add("der", "aze-test")
	d.Get("der")

	/* with CLI */
	// flag.Parse()
	// switch *action {
	// case "list":
	// 	// actionList(d)
	// 	List(d)
	// case "add":
	// 	d.Add(flag.Args())
	// case "define":
	// 	// actionDefine(d, flag.Args())
	// 	d.Define(flag.Args())
	// case "remove":
	// 	// actionRemove(d, flag.Args())
	// 	fmt.Println("Salut les gens")
	// }

	defer d.db.Close()
}

func New() (*DB, error) {
	database, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Ouverture DB %v", database)

	return &DB{
		db: database,
	}, nil
}

func (db DB) CreateBuckets() {
	db.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		fmt.Printf("Create buckets %v", b)
		return nil
	})
}

/*---------------------WITH-CLI----------------------*/

func (db DB) AddAction(args []string) {
	word := args[0]
	definition := args[1]
	created := time.Now()

	e := &Entry{
		Word: word,
		Definition: definition,
		CreatedAt: created,
	}
	fmt.Printf("Initialisation structure Entry: %v", e)

	db.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})
}

func (db DB) DefineAction(args []string) *Entry{
	word := args[0]

	return &Entry{
		Word: word,
	}

}

func (db DB) ViewAction(word string) (Entry, error) {
	var entry Entry

	err := db.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte(word))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
	if err != nil {
		fmt.Printf("View invalid %v", err)
	}
	return entry, err
}

/*---------------------WITHOUT-CLI----------------------*/

func (db DB) Add(word string, definition string) error {
	created := time.Now()
	word = strings.Title(word)
	entry := &Entry{
		Word: word,
		Definition: definition,
		CreatedAt: created,
	}
	fmt.Printf("Initialisation structure Entry: %v", entry)

	var buffer bytes.Buffer 
	enc := gob.NewEncoder(&buffer)
	enc.Encode(entry)

	return db.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		b.Put([]byte(word), buffer.Bytes())
		return nil
	})
}

func (db DB) Get(word string) {
	// var entry Entry

	// var buffer bytes.Buffer 
	// dec := gob.NewDecoder(&buffer)
	// err := dec.Decode(&entry)

	// if err != nil{
	// 	fmt.Printf("Err get: %v\n", err)
	// }
	word = strings.Title(word)

	db.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("Der"))
		fmt.Printf("Valeurs dans le dictionnaire: %v\n", v)
		return nil
	})
}