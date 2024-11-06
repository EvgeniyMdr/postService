package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/EvgeniyMdr/postService/internal/domain/model"
	"github.com/EvgeniyMdr/postService/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (h *Handlers) CreatePost(s services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newPost model.Post
		err := json.NewDecoder(r.Body).Decode(&newPost)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			log.Printf("Error parsing request body: %v", err)
			return
		}

		newPost.ID = uuid.New()

		v := validator.New()

		errValidate := v.Struct(newPost)
		if errValidate != nil {
			for _, err := range errValidate.(validator.ValidationErrors) {
				fmt.Printf("Field '%s' failed validation with tag '%s'\n", err.Field(), err.Tag())
			}
		}

		resp, err := s.CreatePost(r.Context(), newPost)

		if err != nil {
			log.Printf("Error getting post details: %v", err)
			http.Error(w, "Failed to retrieve post details", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			log.Printf("Error encoding response: %v", err)
		}
	}
	//query, err := os.ReadFile("db/sql/create_post.sql")
	//if err != nil {
	//	http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
	//	log.Printf("Error reading SQL file: %v", err)
	//	return
	//}
	//
	//getPostQuery, errGetPostQuery := os.ReadFile("db/sql/get_post_by_id.sql")
	//if errGetPostQuery != nil {
	//	http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
	//	log.Printf("Error reading SQL file: %v", errGetPostQuery)
	//	return
	//}
	//
	//var newPost model.Post
	//err = json.NewDecoder(r.Body).Decode(&newPost)
	//if err != nil {
	//	http.Error(w, "Failed to parse request body", http.StatusBadRequest)
	//	log.Printf("Error parsing request body: %v", err)
	//	return
	//}
	//
	//newPost.ID = uuid.New()
	//
	//v := validator.New()
	//
	//errValidate := v.Struct(newPost)
	//if errValidate != nil {
	//	for _, err := range errValidate.(validator.ValidationErrors) {
	//		fmt.Printf("Field '%s' failed validation with tag '%s'\n", err.Field(), err.Tag())
	//	}
	//}
	//
	//_, err = h.DB.Exec(string(query), newPost.ID, newPost.Title, newPost.Content, newPost.AuthorId, newPost.ImageUrl)
	//if err != nil {
	//	log.Printf("Error inserting post: %v", err)
	//	http.Error(w, "Failed to create post", http.StatusInternalServerError)
	//	return
	//}
	//
	//err = h.DB.QueryRow(string(getPostQuery), newPost.ID).Scan(
	//	&newPost.ID,
	//	&newPost.Title,
	//	&newPost.Content,
	//	&newPost.AuthorId,
	//	&newPost.ImageUrl,
	//	&newPost.CreatedAt,
	//	&newPost.UpdatedAt,
	//)
	//
	//if err != nil {
	//	log.Printf("Error getting post details: %v", err)
	//	http.Error(w, "Failed to retrieve post details", http.StatusInternalServerError)
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusCreated)
	//if err := json.NewEncoder(w).Encode(newPost); err != nil {
	//	http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	//	log.Printf("Error encoding response: %v", err)
	//}
}
