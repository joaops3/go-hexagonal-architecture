package main

import (
	"database/sql"

	db2 "github.com/joaops3/go-exagonal/adapters/db"
	"github.com/joaops3/go-exagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")

	productDbAdapter := db2.NewProductDb(db)

	application.NewProductService(productDbAdapter)
}