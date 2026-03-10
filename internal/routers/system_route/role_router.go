package system_route

import (
	"apis/internal/handler/system"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (r *RoleRouter) Register(public, private *gin.RouterGroup) {
	group := public.Group("/role")
	handler := system.RoleHandler{}
	{
		group.GET("/hello", handler.Hello)
	}
}
