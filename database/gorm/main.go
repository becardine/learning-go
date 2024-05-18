package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           string `gorm:"primary_key"`
	Name         string
	Price        float64
	CategoryID   string
	Category     Category `gorm:"foreignKey:CategoryID"`
	SerialNumber SerialNumber
	gorm.Model   // Add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` to the model
}

type Category struct {
	ID       string `gorm:"primary_key"`
	Name     string
	Products []Product
}

type SerialNumber struct {
	ID        string `gorm:"primary_key"`
	Number    string
	ProductID string
}

func main() {
	dsn := "root:r00t@tcp(localhost:3306)/database?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Create a new product
	p := &Product{Name: "Product 1", Price: 9.99}
	db.Create(p)

	// create batch of products
	products := []Product{
		{Name: "Product 2", Price: 19.99},
		{Name: "Product 3", Price: 29.99},
	}
	db.Create(&products)

	// read a product
	var product Product
	db.First(&product, 1)
	fmt.Println(product)

	// read a product by name
	db.First(&product, "name = ?", "Product 2")

	// read all products
	var productsAll []Product
	db.Find(&productsAll)
	for _, p := range productsAll {
		fmt.Println(p)
	}

	// read all products with limit and offset
	db.Limit(2).Offset(1).Find(&productsAll)

	// read all products with where clause
	db.Where("price > ?", 20).Find(&productsAll)

	// read all products with like clause
	db.Where("name LIKE ?", "%Product%").Find(&productsAll)

	// update a product
	db.Model(&product).Update("Price", 39.99)

	// delete a product
	db.Delete(&product)

	// working with soft delete (deleted_at) need to add gorm.Model to the model
	db.Delete(&product, 1)

	// belongs to relationship
	category := Category{Name: "Category 1"}
	db.Create(&category)
	product = Product{Name: "Product 4", Price: 49.99, CategoryID: category.ID}
	db.Create(&product)

	// select product with category
	db.Preload("Category").Find(&product)

	// has one relationship
	serialNumber := SerialNumber{Number: "123456"}
	db.Create(&serialNumber)
	product = Product{Name: "Product 5", Price: 59.99, SerialNumber: serialNumber}
	db.Create(&product)

	// select product with serial number
	db.Preload("SerialNumber").Find(&product)

	// has many relationship
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, c := range categories {
		fmt.Println(c)
		for _, p := range c.Products {
			fmt.Println(p)
		}
	}
}