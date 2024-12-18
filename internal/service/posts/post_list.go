package posts

import (
	"context"
	"situs-forum/internal/model/posts"

	"github.com/rs/zerolog/log"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error) {
	limit := pageSize
	offset := pageSize * (pageIndex - 1)

	response, err := s.postRepository.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("failed to get all post from service")
		return response, err
	}

	return response, nil
}
