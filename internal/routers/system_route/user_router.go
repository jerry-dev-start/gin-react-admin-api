package system_route

import (
	"apis/internal/handler/system"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (u *UserRouter) Register(public, private *gin.RouterGroup) {
	userHandler := system.NewUserHandler()
	userGroup := public.Group("/user")
	{
		userGroup.POST("", userHandler.SaveUserInfo)
	}
}
