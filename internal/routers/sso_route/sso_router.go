package sso_route

import (
	"apis/internal/handler/system"

	"github.com/gin-gonic/gin"
)

type SsoRouter struct{}

func (s *SsoRouter) Register(public, private *gin.RouterGroup) {
	ssoHandler := system.NewSsoHandler()
	ssoGroup := public.Group("/sso")
	{
		ssoGroup.POST("/login", ssoHandler.SsoLogin)
	}
}
