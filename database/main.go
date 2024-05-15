package main

import "github.com/google/uuid"

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(id int, name string, price float64) *Product {
	return &Product{uuid.New().String(), name, price}
}

func main() {

}
