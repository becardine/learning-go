package main

import (
	"log"
	"net/http"

	"github.com/becardine/learning-go/apis/configs"

	"github.com/becardine/learning-go/apis/internal/entity"
	"github.com/becardine/learning-go/apis/internal/infra/database"
	"github.com/becardine/learning-go/apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use((middleware.Logger))
	// custom middleware to log requests
	// r.Use(LogRequest)

	// Product routes
	r.Route("/products", func(r chi.Router) {
		// Protect the routes with JWT
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		// Use the JWTAuth middleware to check the token validity
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	// User routes
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8080", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
