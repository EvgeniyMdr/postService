package domain

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID        uuid.UUID `json:"id" validate:"required,uuid"`
	Title     string    `json:"title" validate:"required,min=3,max=255"`
	Content   string    `json:"content" validate:"required"`
	AuthorId  uuid.UUID `json:"authorId" validate:"required,uuid"`
	ImageUrl  string    `json:"imageUrl" validate:"omitempty,url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
