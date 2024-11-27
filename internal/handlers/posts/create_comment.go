package posts

import (
	"errors"
	"net/http"
	"situs-forum/internal/model/posts"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()

	var CommentRequest posts.CreateCommentRequest
	if err := c.ShouldBindJSON(&CommentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postID := c.Param("postID")
	postIDStr, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesage": errors.New("post id tidak valid").Error(),
		})
	}

	userID := c.GetInt64("userID")

	err = h.postSvc.CreateComment(ctx, postIDStr, userID, CommentRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success create comment",
	})
}
