package memberships

import (
	"context"
	"situs-forum/internal/configs"
	"situs-forum/internal/model/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
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
