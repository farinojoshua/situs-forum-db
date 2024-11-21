package posts

import (
	"net/http"
	"situs-forum/internal/model/posts"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var PostReq posts.CreatePostRequest
	if err := c.ShouldBindJSON(&PostReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")

	err := h.postSvc.CreatePost(ctx, userID, PostReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "berhasil membuat post"})
}
