package domain

import "github.com/google/uuid"

type UpdatePost struct {
	Title    *string    `json:"title,omitempty"`
	Content  *string    `json:"content,omitempty"`
	AuthorId *uuid.UUID `json:"authorId,omitempty"`
	ImageUrl *string    `json:"imageUrl,omitempty"`
}
