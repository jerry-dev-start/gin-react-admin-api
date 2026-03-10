package orm

import (
	"apis/global"
	"strings"
)

// InitGorm 初始化Gorm
func InitGorm() {
	switch strings.ToLower(global.Config.Server.GetDbType()) {
	case "mysql":
		global.Db = GormMysqlInit()
		return
	default:
		panic("database type not supported")
	}
}
