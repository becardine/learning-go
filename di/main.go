package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	// // create a new repository
	// productRepository := product.NewProductRepository(db)

	// // create a new use case
	// usecase := product.NewProductUseCase(productRepository)

	// create a new use case using wire
	usecase := NewUseCase(db)

	// get a product
	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	// print the product
	fmt.Println(product.Name)
}
