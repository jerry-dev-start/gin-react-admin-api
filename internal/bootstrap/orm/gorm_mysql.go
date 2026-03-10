package orm

import (
	"apis/configs"
	"apis/global"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysqlInit 初始化Mysql数据库
func GormMysqlInit() *gorm.DB {
	//获取到 Mysql 的配置
	m := global.Config.Mysql
	return doMysqlInit(m)
}

func doMysqlInit(m *configs.Mysql) *gorm.DB {
	// 生成DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		m.GetUsername(),
		m.GetPassword(),
		m.GetHost(),
		m.GetPort(),
		m.GetDbName())
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.GetEngine())
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.GetMaxIdleConns())
		sqlDB.SetMaxOpenConns(m.GetMaxIdleConns())
		return db
	}
}
