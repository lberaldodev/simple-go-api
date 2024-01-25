package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"go_api/configs"
	"go_api/internal/entity"
	"go_api/internal/infra/database"
	"go_api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
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

	productDb := database.NewProduct(db)

	userDb := database.NewUser(db)

	productHandler := handlers.NewProductHandler(productDb)

	userHandler := handlers.NewUserHandler(userDb, configs.TokenAuth, configs.JWTExpiresIn)

	r := chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.Create)

		r.Get("/", productHandler.ListProducts)

		r.Get("/{id}", productHandler.GetProduct)

		r.Put("/{id}", productHandler.UpdateProduct)

		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)

	r.Post("/users/login", userHandler.Login)

	http.ListenAndServe(":8080", r)
}
