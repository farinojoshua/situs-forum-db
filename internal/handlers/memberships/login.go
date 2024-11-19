package memberships

import (
	"net/http"
	"situs-forum/internal/model/memberships"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	AccessToken, err := h.membershipSvc.Login(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	response := memberships.LoginResponse{
		AccessToken: AccessToken,
	}

	c.JSON(http.StatusOK, response)
}
