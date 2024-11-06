package http

import (
	"github.com/EvgeniyMdr/postService/internal/http/handlers"
	"github.com/EvgeniyMdr/postService/internal/services"
	"github.com/gorilla/mux"
)

func SetupRouter(service services.Service) *mux.Router {
	h := handlers.Handlers{}

	r := mux.NewRouter()
	r.HandleFunc("/api/posts", h.GetPosts(service)).Methods("GET")

	r.HandleFunc("/api/post/{id}", h.GetPost(service)).Methods("GET")

	r.HandleFunc("/api/post", h.CreatePost(service)).Methods("POST")

	r.HandleFunc("/api/post/{id}", h.DeletePost(service)).Methods("DELETE")

	r.HandleFunc("/api/post/{id}", h.PatchPost(service)).Methods("PATCH")

	r.HandleFunc("/api/post/{id}", h.PutPost(service)).Methods("PUT")
	return r
}
