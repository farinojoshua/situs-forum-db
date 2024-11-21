package posts

import (
	"context"
	"situs-forum/internal/model/posts"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error {
	postHastags := strings.Join(req.PostHashtags, ",")

	model := posts.PostModel{
		UserID:       userID,
		PostTitle:    req.PostTitle,
		PostContent:  req.PostContent,
		PostHashtags: postHastags,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    strconv.FormatInt(userID, 10),
		UpdatedBy:    strconv.FormatInt(userID, 10),
	}

	err := s.postRepository.CreatePost(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to create post")
	}

	return nil
}
