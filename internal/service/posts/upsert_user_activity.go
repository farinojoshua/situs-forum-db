package posts

import (
	"context"
	"errors"
	"situs-forum/internal/model/posts"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, req posts.UserActivityRequest) error {
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   req.IsLiked,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	userActivity, err := s.postRepository.GetUserActivity(ctx, model)
	if err != nil {
		return err
	}

	if userActivity == nil {
		// buat user activity
		if !req.IsLiked {
			return errors.New("anda belum pernah melakukan like sebelumnya")
		}

		err = s.postRepository.CreateUserActivity(ctx, model)

	} else {
		// update user activity
		err = s.postRepository.UpdateUserActivity(ctx, model)
	}
	if err != nil {
		log.Error().Err(err).Msg("error create user activity")
		return err
	}

	return nil
}
