package router

import (
	"apis/global"

	"github.com/gin-gonic/gin"
)

// Router 路由注册器接口。业务方定义结构体并实现 Register，在 public 上挂无需认证的路由，在 private 上挂需认证的路由。
type Router interface {
	Register(public, private *gin.RouterGroup)
}

var (
	routerList     []Router
	authMiddleware gin.HandlerFunc
)

func init() {
	authMiddleware = func(c *gin.Context) { c.Next() }
}

// SetAuthMiddleware 设置需认证路由组使用的认证中间件（如 JWT）。在 Mount 之前调用生效。
func SetAuthMiddleware(mw gin.HandlerFunc) {
	if mw != nil {
		authMiddleware = mw
	}
}

// RegisterRouter 注册一个路由注册器。由 internal/routers 包统一调用，业务只需提供实现 Router 的结构体。
func RegisterRouter(r Router) {
	if r != nil {
		routerList = append(routerList, r)
	}
}

// Mount 将已注册的所有路由挂到 engine 上。
func Mount(engine *gin.Engine) {
	if engine == nil {
		panic("router.Mount: engine is nil")
	}
	prefix := ""
	if global.Config != nil && global.Config.Server != nil {
		prefix = global.Config.Server.GetRouterPrefix()
	}
	public := engine.Group(prefix)
	private := engine.Group(prefix)
	private.Use(authMiddleware)
	for _, r := range routerList {
		r.Register(public, private)
	}
}
