package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/EvgeniyMdr/postService/internal/domain/model"
	"github.com/EvgeniyMdr/postService/internal/domain/requests"
	cError "github.com/EvgeniyMdr/postService/internal/errors"
	"github.com/EvgeniyMdr/postService/internal/repositories/sql_queries"
	"github.com/google/uuid"
	"log"
	"net/http"
	"sync"
	"time"
)

var once sync.Once
var instance Repo

type postRepository struct {
	db *sql.DB
}

func (r postRepository) DeletePost(ctx context.Context, id uuid.UUID) (*bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r postRepository) GetPost(ctx context.Context, id uuid.UUID) (*model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r postRepository) GetPosts(ctx context.Context) ([]*model.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, sql_queries.GetPosts)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			// TODO: Узнать как правильно использовать кастомные ошибки
			return nil, cError.Wrap(errors.New("request timed out. Please try again later"), http.StatusRequestTimeout, "GetPosts")
		}
		log.Printf("Error executing query: %v", err)
		return nil, errors.New("failed to execute query")
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}(rows)

	var posts []*model.Post

	for rows.Next() {
		post := new(model.Post)
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorId, &post.ImageUrl, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, errors.New("failed to scan row")
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return nil, errors.New("error in rows iteration")
	}

	return posts, nil
}

func (r postRepository) CreatePost(ctx context.Context, post model.Post) (*model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r postRepository) PutPost(ctx context.Context, post model.Post) (*model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r postRepository) PatchPost(ctx context.Context, id uuid.UUID, post requests.UpdatePost) (*model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func NewPostRepository(db *sql.DB) Repo {
	once.Do(func() {
		instance = &postRepository{
			db: db,
		}
	})

	return instance
}
