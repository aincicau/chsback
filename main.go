package main

import (
	"chsback/config"
	"chsback/db"
	"chsback/rest"
	"fmt"
	"log"
	"net/http"

	tf "github.com/galeone/tensorflow/tensorflow/go"
	tg "github.com/galeone/tfgo"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	model := tg.LoadModel("model/output/keras", []string{"serve"}, nil)
	fakeInput, _ := tf.NewTensor([1][28][28][1]float32{})
	results := model.Exec([]tf.Output{
		model.Op("StatefulPartitionedCall", 0),
	}, map[tf.Output]*tf.Tensor{
		model.Op("serving_default_inputs_input", 0): fakeInput,
	})
	predictions := results[0]
	fmt.Println(predictions.Value())

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
		r.Get("/history", rest.GetHistory)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
