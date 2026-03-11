package bootstrap

import (
	"apis/global"
	"apis/internal/bootstrap/configs"
	"apis/internal/bootstrap/log"
	"apis/internal/bootstrap/orm"
)

func InitComponent() {
	// 初始化配置文件，读取后放到 global.Config 中
	configs.InitConfig()
	// 初始化日志框架
	log.InitLogger()
	// 初始化 Gorm
	orm.InitGorm()
	global.Log.Info("init component successfully.")
}
