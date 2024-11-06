package services

import (
	"context"
	"github.com/EvgeniyMdr/postService/internal/domain/model"
	"github.com/EvgeniyMdr/postService/internal/domain/requests"
	"github.com/EvgeniyMdr/postService/internal/repositories"
	"github.com/google/uuid"
	"sync"
)

var once sync.Once
var instance Service

type postService struct {
	repo repositories.Repo
}

func (s *postService) DeletePost(ctx context.Context, id uuid.UUID) (*bool, error) {
	resp, err := s.repo.DeletePost(ctx, id)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *postService) GetPost(ctx context.Context, id uuid.UUID) (*model.Post, error) {
	resp, err := s.repo.GetPost(ctx, id)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *postService) CreatePost(ctx context.Context, post model.Post) (*model.Post, error) {
	resp, err := s.repo.CreatePost(ctx, post)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *postService) GetPosts(ctx context.Context) ([]*model.Post, error) {
	resp, err := s.repo.GetPosts(ctx)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *postService) PatchPost(ctx context.Context, id uuid.UUID, post requests.UpdatePost) (*model.Post, error) {
	resp, err := s.repo.PatchPost(ctx, id, post)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *postService) PutPost(ctx context.Context, post model.Post) (*model.Post, error) {
	resp, err := s.repo.PutPost(ctx, post)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewPostService(repo repositories.Repo) Service {
	once.Do(func() {
		instance = &postService{
			repo: repo,
		}
	})

	return instance
}
