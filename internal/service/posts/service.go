package posts

import (
	"context"
	"situs-forum/internal/configs"
	"situs-forum/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error
	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error
}

type service struct {
	cfg            *configs.Config
	postRepository postRepository
}

func NewService(cfg *configs.Config, postRepository postRepository) *service {
	return &service{
		cfg:            cfg,
		postRepository: postRepository,
	}
}
