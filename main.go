package chsback

import (
	"chsback/rest"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/user", rest.CreateUser)
	})
}
