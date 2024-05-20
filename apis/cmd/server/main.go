package main

import (
	"net/http"

	"github.com/becardine/learning-go/apis/configs"

	"github.com/becardine/learning-go/apis/internal/entity"
	"github.com/becardine/learning-go/apis/internal/infra/database"
	"github.com/becardine/learning-go/apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProduct(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8080", nil)
}
