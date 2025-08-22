package main

import (
	"net/http"

	"github.com/brecabral/go-apis/configs"
	"github.com/brecabral/go-apis/internal/entity"
	"github.com/brecabral/go-apis/internal/infra/database"
	"github.com/brecabral/go-apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)

	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8080", r)
}
