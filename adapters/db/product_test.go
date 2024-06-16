package db_test

import (
	"database/sql"
	"log"
	"testing"

	db2 "github.com/joaops3/go-exagonal/adapters/db"
	"github.com/joaops3/go-exagonal/application"
	"github.com/stretchr/testify/require"
)


var Db *sql.DB


func setUp(){
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string,
	);`
	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc", "product test",0,"DISABLED")`
	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}



func TestProductDb_Get(t *testing.T){
	setUp()
	defer Db.Close()

	productDb := db2.NewProductDb(Db)

	product, err := productDb.Get("abc")
	require.Nil(t, err)

	require.Equal(t, "product test", product.GetName())
}

func TestProductDb_Save(t *testing.T){
	setUp()
	defer Db.Close()

	productDb := db2.NewProductDb(Db)

	product := application.NewProduct("teste", 2)

	result, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "teste", result.GetName())

	product.Name = "teste edit"
	result, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, "teste edit", result.GetName())
}

