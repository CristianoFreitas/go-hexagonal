package main

import (
	"database/sql"

	db2 "github.com/cristianofreitas/go-hexagonal/adapters/db"
	"github.com/cristianofreitas/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "dbsqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product", 30)

	productService.Enable(product)
}