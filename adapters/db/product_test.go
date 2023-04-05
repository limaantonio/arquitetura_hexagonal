package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/limaantonio/go-hexagonal/adapters/db"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		id text ,
		name text ,
		price float ,
		status text 
	);`

	stmt, err := Db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	product := `INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?);`

	stmt, err := Db.Prepare(product)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec("1", "product 1", 1.0, "disabled")
}

func TestProductDb_Get(t *testing.T) {
	setup()
	//defer é um statment que executa uma função quando a função que o contém termina
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("1")
	require.Nil(t, err)
	require.Equal(t, "product 1", product.GetName())
	require.Equal(t, 1.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())

}
