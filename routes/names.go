package routes

import (
	"encoding/json"
	"net/http"

	"github.com/cmar82_nike/go-chi-api/db"
	"github.com/cmar82_nike/go-chi-api/models"
	"github.com/go-chi/chi"
)

type NamesResource struct{}

func (nr NamesResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", nr.Get) // GET /names - Read a single name
	return r
}

// Request Handler - GET /names - Read a single name
func (nr NamesResource) Get(w http.ResponseWriter, r *http.Request) {
	// name := db.RandomName()
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(name))
	name := db.RandomName()
	resp := &models.NameResponse{Name: name}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
