package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"github.com/cmar82_nike/go-chi-api/db"
	"github.com/cmar82_nike/go-chi-api/routes"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	port := "8000"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello World!"))
	})

	r.Mount("/posts", routes.PostsResource{}.Routes())
	r.Mount("/names", routes.NamesResource{}.Routes())

	log.Fatal(http.ListenAndServe(":"+port, r))
}
