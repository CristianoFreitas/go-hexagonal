package main

import (
	"database/sql"
	"log"

	db2 "github.com/cristianofreitas/go-hexagonal/adapters/db"
	"github.com/cristianofreitas/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, err := productService.Create("Product", 30)

	if err != nil {
		log.Fatal(err.Error())
	} else {

		productService.Enable(product)
	}

}
