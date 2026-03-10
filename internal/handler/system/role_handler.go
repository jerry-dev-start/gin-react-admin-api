package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
}

func (h *RoleHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}
