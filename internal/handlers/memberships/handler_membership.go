package memberships

import (
	"context"
	"situs-forum/internal/model/memberships"

	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SignUp(ctx context.Context, request memberships.SignUpRequest) error
}

type Handler struct {
	*gin.Engine

	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")
	route.GET("/ping", h.Ping)
	route.POST("/sign-up", h.SignUp)
}
