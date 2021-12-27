package main

import (
	"chsback/config"
	"chsback/db"
	"chsback/rest"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatal("Not able to create config")
	}

	db.InitDatabase()

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Route("/", func(r chi.Router) {
		r.Post("/user", rest.CreateUser)
		r.Post("/login", rest.Login)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
