package main

import (
	
)

func (d *Entry) Add() {
	
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