package handlers

import (
	"encoding/json"
	"github.com/EvgeniyMdr/postService/internal/services"
	"log"
	"net/http"
)

func (h *Handlers) GetPosts(s services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := s.GetPosts(r.Context())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error executing query: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(posts); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			log.Printf("Error encoding response: %v", err)
		}
	}
	//query, err := os.ReadFile("db/sql/get_all_posts.sql")
	//if err != nil {
	//	http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
	//	log.Printf("Error reading SQL file: %v", err)
	//	return
	//}
	//
	//rows, err := h.DB.Query(string(query))
	//if err != nil {
	//	http.Error(w, "Failed to execute query", http.StatusInternalServerError)
	//	log.Printf("Error executing query: %v", err)
	//	return
	//}
	//
	//defer func(rows *sql.Rows) {
	//	err := rows.Close()
	//	if err != nil {
	//		log.Printf("Error closing rows: %v", err)
	//	}
	//}(rows)
	//
	//var posts []model.Post
	//
	//for rows.Next() {
	//	var post model.Post
	//	err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorId, &post.ImageUrl, &post.CreatedAt, &post.UpdatedAt)
	//	if err != nil {
	//		http.Error(w, "Failed to scan row", http.StatusInternalServerError)
	//		log.Printf("Error scanning row: %v", err)
	//		return
	//	}
	//	posts = append(posts, post)
	//}
	//
	//if err = rows.Err(); err != nil {
	//	http.Error(w, "Error in rows iteration", http.StatusInternalServerError)
	//	log.Printf("Error iterating rows: %v", err)
	//	return
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//if err := json.NewEncoder(w).Encode(posts); err != nil {
	//	http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	//	log.Printf("Error encoding response: %v", err)
	//}
}
