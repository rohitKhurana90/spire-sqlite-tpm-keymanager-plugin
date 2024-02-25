package main

import (
	_ "github.com/mattn/go-sqlite3"
)

type signingKeys struct {
	id          int
	keyValue    string
	typeValue   string
	randomValue string
}
