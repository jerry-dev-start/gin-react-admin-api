package global

import (
	"apis/configs"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *configs.Config
	Log    *zap.SugaredLogger
	Db     *gorm.DB
)
