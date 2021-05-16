package main

import (
	
)

func (store *dbStore) Add(title string, definition string) {
	entry := &Entry{
		Title: title,
		Definition: definition,
	}
	tx := store.db.MustBegin()
	tx.NamedExec("INSERT INTO entry (title, definition) VALUES (:title, :definition)", entry)
	tx.Commit()
}

func (store *dbStore) AddMany(title string, definition string) {
	// TODO add title and definition with num suffix

	entry := &Entry{
		Title: title,
		Definition: definition,
	}
	for i := 0; i < 10; i++ {
		tx := store.db.MustBegin()
		tx.NamedExec("INSERT INTO entry (title, definition) VALUES (:title, :definition)", entry)
		tx.Commit()
	}
}

func (store *dbStore) Get(title string) (Entry, error){
	// show entity who is selected by word
	entry := Entry{}
    err := store.db.Get(&entry, "SELECT * FROM entry WHERE title = ?", title)
	if err != nil {
		return entry, err
	}
    return entry, err
}

func (store *dbStore) Update(title string, definition string) {
	// update definition with word for selecting
	tx := store.db.MustBegin()
	tx.MustExec("UPDATE entry SET definition = ? WHERE title = ?", definition, title)
	tx.Commit()
}

func (store *dbStore) List() ([]Entry, error) {
	var entry []Entry
	err := store.db.Select(&entry, "SELECT * FROM entry")
	if err != nil {
		return entry, err
	}
	return entry, nil
}

func (store *dbStore) Remove(title string) {
	// TODO if database is empty alert message
	tx := store.db.MustBegin()
	tx.MustExec("DELETE FROM entry WHERE title = ?", title)
	tx.Commit()
}

func (store *dbStore) RemoveAll() {
	tx := store.db.MustBegin()
	tx.MustExec("DELETE FROM entry")
	tx.Commit()
}
