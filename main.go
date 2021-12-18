package chsback

import (
	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/history", ...)
	})
}
