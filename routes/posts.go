package routes

import (
	"context"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

type PostsResource struct{}

func (pr PostsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", pr.List)    // GET /posts - Read a list of posts.
	r.Post("/", pr.Create) // POST /posts - Create a new post.

	r.Route("/{id}", func(r chi.Router) {
		r.Use(PostCtx)
		r.Get("/", pr.Get)       // GET /posts/{id} - Read a single post by :id.
		r.Put("/", pr.Update)    // PUT /posts/{id} - Update a single post by :id.
		r.Delete("/", pr.Delete) // DELETE /posts/{id} - Delete a single post by :id.
	})

	return r
}

// Request Handler - GET /posts - Read a list of posts.
func (pr PostsResource) List(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Request Handler - POST /posts - Create a new post.
func (pr PostsResource) Create(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Request Handler - GET /posts/{id} - Read a single post by :id.
func (pr PostsResource) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Request Handler - PUT /posts/{id} - Update a single post by :id.
func (pr PostsResource) Update(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	client := &http.Client{}

	req, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/posts/"+id, r.Body)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Request Handler - DELETE /posts/{id} - Delete a single post by :id.
func (pr PostsResource) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/"+id, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
