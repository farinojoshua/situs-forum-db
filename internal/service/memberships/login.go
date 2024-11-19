package memberships

import (
	"context"
	"errors"
	"situs-forum/internal/model/memberships"
	"situs-forum/pkg/jwt"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, request memberships.LoginRequest) (string, error) {
	user, err := s.membershipsRepo.GetUser(ctx, request.Email)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("email not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", err
	}

	token, err := jwt.CreateToken(user.UserID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}

	return token, nil
}
