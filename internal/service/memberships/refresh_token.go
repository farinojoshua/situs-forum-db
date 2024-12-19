package memberships

import (
	"context"
	"errors"
	"situs-forum/internal/model/memberships"
	"situs-forum/pkg/jwt"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipsRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get refresh token from db")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has been expired")
	}

	// token in db is not match with the request token
	if existingRefreshToken.RefreshToken != request.Token {
		return "", errors.New("refresh token is valid")
	}

	user, err := s.membershipsRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("user not exists")
	}

	token, err := jwt.CreateToken(user.UserID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}

	return token, nil
}
