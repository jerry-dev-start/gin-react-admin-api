// Package routers 统一注册各业务模块的路由器。新增模块时在此文件追加一行 RegisterRouter 即可，无需改 main、无需 init。
package routers

import (
	bootRouter "apis/internal/bootstrap/router"
	"apis/internal/routers/sso_route"
	"apis/internal/routers/system_route"
)

func init() {
	bootRouter.RegisterRouter(&system_route.RoleRouter{})
	//单点登陆相关路由
	bootRouter.RegisterRouter(&sso_route.SsoRouter{})
}
