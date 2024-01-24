package main

import (
	"github.com/go-chi/chi"
	"go_api/configs"
	"go_api/internal/entity"
	"go_api/internal/infra/database"
	"go_api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
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

	productDb := database.NewProduct(db)

	productHandler := handlers.NewProductHandler(productDb)

	r := chi.NewRouter()

	r.Post("/products", productHandler.Create)

	r.Get("/product/{id}", productHandler.GetProduct)

	r.Put("/product/{id}", productHandler.UpdateProduct)

	r.Delete("/product/{id}", productHandler.DeleteProduct)

	http.ListenAndServe(":8080", r)
}
