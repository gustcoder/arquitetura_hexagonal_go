package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
	//productDB.Exec("CREATE TABLE products (id string, name string, price float, status string)")
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products("id" string, "name" string, "price" float, "status" string);`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES ("123", "PlayStation5", 3500.00, "enabled");`
	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDB := db.NewProductDb(Db)

	product, err := productDB.Get("123")
	require.Nil(t, err)
	require.Equal(t, "PlayStation5", product.GetName())
	require.Equal(t, 3500.00, product.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())
}
