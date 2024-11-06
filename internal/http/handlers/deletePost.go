package handlers

import (
	"github.com/EvgeniyMdr/postService/internal/services"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (h *Handlers) DeletePost(s services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		uuidValue, err := uuid.Parse(id)

		if err != nil {
			http.Error(w, "Invalid UUID format:", http.StatusInternalServerError)
			log.Printf("Invalid UUID format: %v", err)
		}

		resp, err := s.DeletePost(r.Context(), uuidValue)

		if err != nil {
			http.Error(w, "Failed to execute delete query", http.StatusInternalServerError)
			log.Printf("Error executing delete query: %v", err)
			return
		}

		if resp != nil {
			w.WriteHeader(http.StatusOK)
		}
	}
	//query, err := os.ReadFile("db/sql/delete_post.sql")
	//if err != nil {
	//	http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
	//	log.Printf("Error reading SQL file: %v", err)
	//	return
	//}
	//
	//vars := mux.Vars(r)
	//id := vars["id"]
	//
	//result, err := h.DB.Exec(string(query), id)
	//if err != nil {
	//	http.Error(w, "Failed to execute delete query", http.StatusInternalServerError)
	//	log.Printf("Error executing delete query: %v", err)
	//	return
	//}
	//
	//rowsAffected, err := result.RowsAffected()
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
	//w.WriteHeader(http.StatusOK)
}
