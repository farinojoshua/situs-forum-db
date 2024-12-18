package posts

import (
	"context"
	"situs-forum/internal/model/posts"

	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepository.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id to db")
		return nil, err
	}

	likeCount, err := s.postRepository.CountLikeByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get count like by id to db")
		return nil, err
	}

	comments, err := s.postRepository.GetCommentByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get comments by id to db")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			PostID:       postDetail.PostID,
			UserID:       postDetail.UserID,
			Username:     postDetail.Username,
			PostTitle:    postDetail.PostTitle,
			PostContent:  postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked:      postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments:  comments,
	}, nil
}
