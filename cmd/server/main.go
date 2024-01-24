package main

import (
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

	http.HandleFunc("/products", productHandler.Create)

	http.ListenAndServe(":8080", nil)
}
