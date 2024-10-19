package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/EvgeniyMdr/postService/internal/domain"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
)

func (h *Handlers) GetPost(w http.ResponseWriter, r *http.Request) {
	query, err := os.ReadFile("db/sql/get_post_by_id.sql")
	if err != nil {
		http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
		log.Printf("Error reading SQL file: %v", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var post domain.Post
	err = h.DB.QueryRow(string(query), id).Scan(&post.ID, &post.Title, &post.Content, &post.AuthorId, &post.ImageUrl, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			log.Printf("Error scanning row: %v", err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
	}
}
