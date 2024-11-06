package handlers

import (
	"encoding/json"
	"github.com/EvgeniyMdr/postService/internal/services"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (h *Handlers) GetPost(s services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		uuidValue, err := uuid.Parse(id)

		if err != nil {
			http.Error(w, "Invalid UUID format:", http.StatusInternalServerError)
			log.Printf("Invalid UUID format: %v", err)
		}
		resp, err := s.GetPost(r.Context(), uuidValue)

		if err != nil {
			http.Error(w, "Failed to execute query", http.StatusInternalServerError)
			log.Printf("Error executing query: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			log.Printf("Error encoding response: %v", err)
		}
	}
	//query, err := os.ReadFile("db/sql/get_post_by_id.sql")
	//if err != nil {
	//	http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
	//	log.Printf("Error reading SQL file: %v", err)
	//	return
	//}
	//
	//vars := mux.Vars(r)
	//id := vars["id"]
	//
	//var post model.Post
	//err = h.DB.QueryRow(string(query), id).Scan(&post.ID, &post.Title, &post.Content, &post.AuthorId, &post.ImageUrl, &post.CreatedAt, &post.UpdatedAt)
	//
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		http.Error(w, "Post not found", http.StatusNotFound)
	//	} else {
	//		http.Error(w, "Failed to scan row", http.StatusInternalServerError)
	//		log.Printf("Error scanning row: %v", err)
	//	}
	//	return
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//if err := json.NewEncoder(w).Encode(post); err != nil {
	//	http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	//	log.Printf("Error encoding response: %v", err)
	//}
}
