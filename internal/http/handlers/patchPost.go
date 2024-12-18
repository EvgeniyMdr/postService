package handlers

import (
	"encoding/json"
	"github.com/EvgeniyMdr/postService/internal/domain/requests"
	"github.com/EvgeniyMdr/postService/internal/services"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (h *Handlers) PatchPost(s services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		uuidValue, err := uuid.Parse(id)

		if err != nil {
			http.Error(w, "Invalid UUID format:", http.StatusInternalServerError)
			log.Printf("Invalid UUID format: %v", err)
		}

		var updatePost requests.UpdatePost
		err = json.NewDecoder(r.Body).Decode(&updatePost)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			log.Printf("Error parsing request body: %v", err)
			return
		}

		resp, err := s.PatchPost(r.Context(), uuidValue, updatePost)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			log.Printf("Error encoding response: %v", err)
		}
	}
	//query, err := os.ReadFile("db/sql/patch_post.sql")
	//if err != nil {
	//	http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
	//	log.Printf("Error reading SQL file: %v", err)
	//	return
	//}
	//
	//vars := mux.Vars(r)
	//id := vars["id"]
	//
	//var updatePost requests.UpdatePost
	//err = json.NewDecoder(r.Body).Decode(&updatePost)
	//if err != nil {
	//	http.Error(w, "Failed to parse request body", http.StatusBadRequest)
	//	log.Printf("Error parsing request body: %v", err)
	//	return
	//}
	//
	//result, err := h.DB.Exec(string(query),
	//	updatePost.Title,
	//	updatePost.Content,
	//	updatePost.AuthorId,
	//	updatePost.ImageUrl,
	//	id,
	//)
	//
	//if err != nil {
	//	http.Error(w, "Failed to execute patch query", http.StatusInternalServerError)
	//	log.Printf("Error executing patch query: %v", err)
	//	return
	//}
	//
	//rowsAffected, err := result.RowsAffected()
	//
	//if err != nil {
	//	http.Error(w, "Error checking affected rows", http.StatusInternalServerError)
	//	log.Printf("Error checking affected rows: %v", err)
	//	return
	//}
	//
	//if rowsAffected == 0 {
	//	http.Error(w, "Post not found", http.StatusNotFound)
	//	return
	//}
	//
	//getPostQuery, err := os.ReadFile("db/sql/get_post_by_id.sql")
	//if err != nil {
	//	http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
	//	log.Printf("Error reading SQL file: %v", err)
	//	return
	//}
	//
	//var post model.Post
	//err = h.DB.QueryRow(string(getPostQuery), id).Scan(
	//	&post.ID,
	//	&post.Title,
	//	&post.Content,
	//	&post.AuthorId,
	//	&post.ImageUrl,
	//	&post.CreatedAt,
	//	&post.UpdatedAt,
	//)
	//
	//if err != nil {
	//	http.Error(w, "Failed to retrieve updated post", http.StatusInternalServerError)
	//	log.Printf("Error retrieving updated post: %v", err)
	//	return
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//if err := json.NewEncoder(w).Encode(post); err != nil {
	//	http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	//	log.Printf("Error encoding response: %v", err)
	//}
}
