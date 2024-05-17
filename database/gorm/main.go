package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    string `gorm:"primary_key"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:r00t@tcp(localhost:3306)/database?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// Create a new product
	p := &Product{Name: "Product 1", Price: 9.99}
	db.Create(p)

	// create batch of products
	products := []Product{
		{Name: "Product 2", Price: 19.99},
		{Name: "Product 3", Price: 29.99},
	}
	db.Create(&products)
}
