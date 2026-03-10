package system

import (
	"github.com/gin-gonic/gin"
)

// SystemRouter 实现 bootstrap/router 的 Router 接口，注册 system_route 模块的路由。
type SystemRouter struct{}

// Register 在 public/private 上注册路由，由框架在 Mount 时调用。
func (r *SystemRouter) Register(public, private *gin.RouterGroup) {
	roleHandler := &RoleHandler{}
	public.GET("/hello", roleHandler.Hello)
}
