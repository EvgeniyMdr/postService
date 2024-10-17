package http

import (
	"database/sql"
	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) *mux.Router {
	h := &Handlers{DB: db}

	r := mux.NewRouter()
	r.HandleFunc("/api/posts", h.GetAllPosts).Methods("GET")

	r.HandleFunc("/api/post/{id}", h.GetPost).Methods("GET")

	r.HandleFunc("/api/post", h.CreatePost).Methods("POST")

	r.HandleFunc("/api/post/{id}", h.DeletePost).Methods("DELETE")

	r.HandleFunc("/api/post/{id}", h.PatchPost).Methods("PATCH")

	r.HandleFunc("/api/post/{id}", h.PutPost).Methods("PUT")
	return r
}
