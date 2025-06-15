package main

import (
	"database/sql"

	adapter "github.com/gustcoder/arquitetura_hexagonal_go/adapters/db"
	"github.com/gustcoder/arquitetura_hexagonal_go/application"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")

	productDbAdapter := adapter.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("CRF 250F", 25000)

	productService.Enable(product)
}
