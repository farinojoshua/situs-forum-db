package posts

import (
	"context"
	"situs-forum/internal/model/posts"
)

func (r *repository) CreateComment(ctx context.Context, model posts.CommentModel) error {
	query := `INSERT INTO comments (post_id, user_id, comment_content, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.CommentContent, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetCommentByPostID(ctx context.Context, postID int64) ([]posts.Comment, error) {
	query := `SELECT c.comment_id, c.user_id, c.comment_content, u.username FROM comments c JOIN users u ON c.user_id = u.user_id WHERE c.post_id = ?`

	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := make([]posts.Comment, 0)
	for rows.Next() {
		var (
			comment  posts.CommentModel
			username string
		)

		err = rows.Scan(&comment.CommentID, &comment.UserID, &comment.CommentContent, &username)
		if err != nil {
			return nil, err
		}

		response = append(response, posts.Comment{
			CommentID:      comment.CommentID,
			UserID:         comment.UserID,
			CommentContent: comment.CommentContent,
			Username:       username,
		})
	}

	return response, nil
}
