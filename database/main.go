package main

import (
	"database/sql"

	// Import the mysql driver package so that we can use it to connect to the database
	// The blank identifier is used to import a package solely for its side-effects (initialization)
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{uuid.New().String(), name, price}
}

func main() {
	// Open the connection to the database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/database")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Create a new product
	p := NewProduct("Product 1", 9.99)
	err = insertProduct(db, p)
	if err != nil {
		panic(err)
	}

	// Update the product
	p.Price = 19.99
	err = updateProduct(db, p)
	if err != nil {
		panic(err)
	}

}

func insertProduct(db *sql.DB, p *Product) error {
	// Prepare the statement to insert a new product
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(p.ID, p.Name, p.Price)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, p *Product) error {
	// Prepare the statement to update a product
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(p.Name, p.Price, p.ID)
	if err != nil {
		return err
	}

	return nil
}
