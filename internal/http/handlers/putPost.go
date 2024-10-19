package handlers

import (
	"encoding/json"
	"github.com/EvgeniyMdr/postService/internal/domain"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func (h *Handlers) PutPost(w http.ResponseWriter, r *http.Request) {
	query, err := os.ReadFile("db/sql/put_post.sql")
	if err != nil {
		http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
		log.Printf("Error reading SQL file: %v", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var updatedPost domain.Post
	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}

	if updatedPost.Title == "" || updatedPost.Content == "" || updatedPost.AuthorId == uuid.Nil || updatedPost.ImageUrl == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	result, err := h.DB.Exec(string(query),
		updatedPost.Title,
		updatedPost.Content,
		updatedPost.AuthorId,
		updatedPost.ImageUrl,
		id,
	)

	if err != nil {
		http.Error(w, "Failed to execute put query", http.StatusInternalServerError)
		log.Printf("Error executing put query: %v", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error checking affected rows", http.StatusInternalServerError)
		log.Printf("Error checking affected rows: %v", err)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	getPostQuery, err := os.ReadFile("db/sql/get_post_by_id.sql")
	if err != nil {
		http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
		log.Printf("Error reading SQL file: %v", err)
		return
	}

	var post domain.Post
	err = h.DB.QueryRow(string(getPostQuery), id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.AuthorId,
		&post.ImageUrl,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		http.Error(w, "Failed to retrieve updated post", http.StatusInternalServerError)
		log.Printf("Error retrieving updated post: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
	}
}
