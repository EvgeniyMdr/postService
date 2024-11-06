package services

import (
	"context"
	"github.com/EvgeniyMdr/postService/internal/domain/model"
	"github.com/EvgeniyMdr/postService/internal/domain/requests"
	"github.com/google/uuid"
)

type Service interface {
	GetPost(ctx context.Context, id uuid.UUID) (*model.Post, error)
	CreatePost(ctx context.Context, post model.Post) (*model.Post, error)
	GetPosts(ctx context.Context) ([]*model.Post, error)
	PatchPost(ctx context.Context, id uuid.UUID, post requests.UpdatePost) (*model.Post, error)
	PutPost(ctx context.Context, post model.Post) (*model.Post, error)
	DeletePost(ctx context.Context, id uuid.UUID) (*bool, error)
}
