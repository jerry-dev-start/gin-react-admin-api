// Package router 提供按「需认证 / 无需认证」批量注册路由的能力。
//
// 使用方式：
//
//  1. 在业务包中定义实现 Router 接口的结构体，在 Register 里往 public（无需认证）/ private（需认证）上挂路由：
//
//     type UserRouter struct{}
//     func (r *UserRouter) Register(public, private *gin.RouterGroup) {
//     public.POST("/login", ...)
//     private.GET("/user/info", ...)
//     }
//
//  2. 在 internal/routers/register.go 的 init 中追加一行，注册该路由器：
//
//     bootRouter.RegisterRouter(&user.UserRouter{})
//
//  3. main 中已空白导入 internal/routers，无需再改 main；Mount 前可设置认证中间件：
//
//     router.SetAuthMiddleware(middleware.JWTAuth())
//     router.Mount(s.Engine)
package router
