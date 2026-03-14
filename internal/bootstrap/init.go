package bootstrap

import (
	"apis/global"
	"apis/internal/bootstrap/configs"
	"apis/internal/bootstrap/log"
	"apis/internal/bootstrap/orm"
	"apis/internal/bootstrap/snowflake"
	"apis/internal/bootstrap/sso"
)

func InitComponent() {
	// 初始化配置文件，读取后放到 global.Config 中
	configs.InitConfig()
	// 初始化日志框架
	log.InitLogger()
	// 初始化 Gorm
	orm.InitGorm()

	//初始化雪花算法
	if err := snowflake.InitSnowflake("2001-01-01", 1); err != nil {
		panic(err)
	}

	//初始化单点登录
	if global.Config.Server.GetOpenSso() {
		sso.InitCasDoor(global.Config.CasdoorConfig)
	}
	global.Log.Info("init component successfully.")
}
