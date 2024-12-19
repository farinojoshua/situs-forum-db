package memberships

import (
	"context"
	"errors"
	"situs-forum/internal/model/memberships"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipsRepo.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("username or email already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	model := memberships.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(pass),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	return s.membershipsRepo.CreateUser(ctx, model)
}
