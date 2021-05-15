package main

import (
	
)

func (store *dbStore) Add(title string, definition string) {
	// ./cli -action add Kotlin "A great contender to Java"
	entry := &Entry{
		Title: title,
		Definition: definition,
	}
	tx := store.db.MustBegin()
	tx.NamedExec("INSERT INTO entry (title, definition) VALUES (:title, :definition)", entry)
	tx.Commit()
}

func (d *Entry) Get() {
	
}

func (d *Entry) Define() {
	
}

func (store *dbStore) List() ([]Entry, error) {
	var entry []Entry
	err := store.db.Select(&entry, "SELECT * FROM entry")
	if err != nil {
		return entry, err
	}
	return entry, nil
}