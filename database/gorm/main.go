package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         string `gorm:"primary_key"`
	Name       string
	Price      float64
	CategoryID string
	Category   Category `gorm:"foreignKey:CategoryID"`
	// Categories   []Category `gorm:"many2many:product_categories;"` // many to many relationship
	SerialNumber SerialNumber
	gorm.Model   // Add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` to the model
}

type Category struct {
	ID       string `gorm:"primary_key"`
	Name     string
	Products []Product // `gorm:"many2many:product_categories;"`  many to many relationship
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
	// Products.SerialNumber will be populated with the serial number of each product in the category when it's loaded
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	// alternatively, you can use the following code to load the products and serial numbers
	// err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, c := range categories {
		fmt.Println(c)
		for _, p := range c.Products {
			fmt.Println(p)
		}
	}

	// many to many relationship
	// db.Model(&product).Association("Categories").Append(&category)
	// db.Model(&product).Association("Categories").Delete(&category)
	// db.Model(&product).Association("Categories").Find(&categories)
	// db.Model(&product).Association("Categories").Count()
	// db.Model(&product).Association("Categories").Clear()
	// db.Model(&product).Association("Categories").Replace(&category)
	// db.Create(&Product{
	// 	Name: "Product 6",
	// 	Price: 69.99,
	// 	Categories: []Category{
	// 		{Name: "Category 2"},
	// 		{Name: "Category 3"},
	// 	},
	// })

}
