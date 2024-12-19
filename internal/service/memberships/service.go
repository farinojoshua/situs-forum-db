package memberships

import (
	"context"
	"situs-forum/internal/configs"
	"situs-forum/internal/model/memberships"
	"time"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error)
	InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error
}
type service struct {
	cfg             *configs.Config
	membershipsRepo membershipRepository
}

func NewService(cfg *configs.Config, membershipRepo membershipRepository) *service {
	return &service{
		cfg:             cfg,
		membershipsRepo: membershipRepo,
	}
}
