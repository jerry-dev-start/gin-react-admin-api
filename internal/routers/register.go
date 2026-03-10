// Package routers 统一注册各业务模块的路由器。新增模块时在此文件追加一行 RegisterRouter 即可，无需改 main、无需 init。
package routers

import (
	bootRouter "apis/internal/bootstrap/router"
	"apis/internal/routers/system_route"
)

func init() {
	bootRouter.RegisterRouter(&system_route.RoleRouter{})
}
