package main

import (
	"chsback/config"
	"chsback/db"
	"chsback/rest"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatal("Not able to create config")
	}

	db.InitDatabase()

	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Post("/user", rest.CreateUser)
		r.Post("/login", rest.Login)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
