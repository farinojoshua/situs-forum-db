package posts

import (
	"context"
	"situs-forum/internal/middleware"
	"situs-forum/internal/model/posts"

	"github.com/gin-gonic/gin"
)

type PostService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
}

type Handler struct {
	*gin.Engine

	postSvc PostService
}

func NewHandler(api *gin.Engine, postSvc PostService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create", h.CreatePost)
}