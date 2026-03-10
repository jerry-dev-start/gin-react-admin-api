package main

import (
	"apis/global"
	"apis/internal/bootstrap"
	"apis/internal/bootstrap/gin"
	bootRouter "apis/internal/bootstrap/router"

	_ "apis/internal/routers" // 拉取路由汇聚包，自动注册各业务 Router
)

func main() {
	// 初始化各个组件
	bootstrap.InitComponent()
	// 创建 Gin 服务并挂载路由（需认证 / 无需认证 由各业务通过 router.Register 注册）
	s := gin.NewServer(global.Config)
	bootRouter.Mount(s.Engine)
	s.StartWeb()
}
