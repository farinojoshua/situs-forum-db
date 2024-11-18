package memberships

import (
	"context"
	"situs-forum/internal/model/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}
type service struct {
	membershipsRepo membershipRepository
}

func NewService(membershipRepo membershipRepository) *service {
	return &service{
		membershipsRepo: membershipRepo,
	}
}
