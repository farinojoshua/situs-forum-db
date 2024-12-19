package memberships

import (
	"context"
	"errors"
	"situs-forum/internal/model/memberships"
	"situs-forum/pkg/jwt"
	tokenUtil "situs-forum/pkg/token"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, request memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipsRepo.GetUser(ctx, request.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("email not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", "", err
	}

	token, err := jwt.CreateToken(user.UserID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", "", err
	}

	existingRefreshToken, err := s.membershipsRepo.GetRefreshToken(ctx, user.UserID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get latest refresh token from db")
		return "", "", err
	}

	if existingRefreshToken != nil {
		return token, existingRefreshToken.RefreshToken, nil
	}

	refreshToken := tokenUtil.GenerateRefreshToken()
	if refreshToken == "" {
		return token, "", errors.New("failed to generate refresh token")
	}

	err = s.membershipsRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserID:       user.UserID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(10 * 24 * time.Hour),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    strconv.Itoa(int(user.UserID)),
		UpdatedBy:    strconv.Itoa(int(user.UserID)),
	})
	if err != nil {
		log.Error().Err(err).Msg("error inserting token to db")
		return token, refreshToken, err
	}

	return token, refreshToken, nil
}
