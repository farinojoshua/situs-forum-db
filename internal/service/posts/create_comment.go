package posts

import (
	"context"
	"situs-forum/internal/model/posts"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) CreateComment(ctx context.Context, postID, userID int64, req posts.CreateCommentRequest) error {
	model := posts.CommentModel{
		PostID:         postID,
		UserID:         userID,
		CommentContent: req.CommentContent,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		CreatedBy:      strconv.FormatInt(userID, 10),
		UpdatedBy:      strconv.FormatInt(userID, 10),
	}

	err := s.postRepository.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create comment to repository")
		return err
	}

	return nil
}
