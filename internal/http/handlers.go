package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/EvgeniyMdr/postService/internal/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
)

type Handlers struct {
	DB *sql.DB
}

func (h *Handlers) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	query, err := os.ReadFile("db/sql/get_all_posts.sql")
	if err != nil {
		http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
		log.Printf("Error reading SQL file: %v", err)
		return
	}

	rows, err := h.DB.Query(string(query))
	if err != nil {
		http.Error(w, "Failed to execute query", http.StatusInternalServerError)
		log.Printf("Error executing query: %v", err)
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}(rows)

	var posts []domain.Post

	for rows.Next() {
		var post domain.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorId, &post.ImageUrl, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			log.Printf("Error scanning row: %v", err)
			return
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Error in rows iteration", http.StatusInternalServerError)
		log.Printf("Error iterating rows: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
	}
}

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

func (h *Handlers) CreatePost(w http.ResponseWriter, r *http.Request) {
	query, err := os.ReadFile("db/sql/create_post.sql")
	if err != nil {
		http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
		log.Printf("Error reading SQL file: %v", err)
		return
	}

	getPostQuery, errGetPostQuery := os.ReadFile("db/sql/get_post_by_id.sql")
	if errGetPostQuery != nil {
		http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
		log.Printf("Error reading SQL file: %v", errGetPostQuery)
		return
	}

	var newPost domain.Post
	err = json.NewDecoder(r.Body).Decode(&newPost)
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

	_, err = h.DB.Exec(string(query), newPost.ID, newPost.Title, newPost.Content, newPost.AuthorId, newPost.ImageUrl)
	if err != nil {
		log.Printf("Error inserting post: %v", err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	err = h.DB.QueryRow(string(getPostQuery), newPost.ID).Scan(
		&newPost.ID,
		&newPost.Title,
		&newPost.Content,
		&newPost.AuthorId,
		&newPost.ImageUrl,
		&newPost.CreatedAt,
		&newPost.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error getting post details: %v", err)
		http.Error(w, "Failed to retrieve post details", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newPost); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
	}
}

func (h *Handlers) DeletePost(w http.ResponseWriter, r *http.Request) {
	query, err := os.ReadFile("db/sql/delete_post.sql")
	if err != nil {
		http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
		log.Printf("Error reading SQL file: %v", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	result, err := h.DB.Exec(string(query), id)
	if err != nil {
		http.Error(w, "Failed to execute delete query", http.StatusInternalServerError)
		log.Printf("Error executing delete query: %v", err)
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

	w.WriteHeader(http.StatusOK)
}

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

func (h *Handlers) PatchPost(w http.ResponseWriter, r *http.Request) {
	query, err := os.ReadFile("db/sql/patch_post.sql")
	if err != nil {
		http.Error(w, "Failed to read SQL file", http.StatusInternalServerError)
		log.Printf("Error reading SQL file: %v", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var updatePost domain.UpdatePost
	err = json.NewDecoder(r.Body).Decode(&updatePost)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}

	result, err := h.DB.Exec(string(query),
		updatePost.Title,
		updatePost.Content,
		updatePost.AuthorId,
		updatePost.ImageUrl,
		id,
	)

	if err != nil {
		http.Error(w, "Failed to execute patch query", http.StatusInternalServerError)
		log.Printf("Error executing patch query: %v", err)
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
