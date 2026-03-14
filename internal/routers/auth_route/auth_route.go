package auth_route

import (
	"apis/internal/handler/system"

	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
}

func (a *AuthRoute) Register(public, private *gin.RouterGroup) {
	authHandler := system.NewAuthHandler()
	publicGroup := public.Group("/auth")
	{
		publicGroup.POST("/login", authHandler.Login)
	}
}
