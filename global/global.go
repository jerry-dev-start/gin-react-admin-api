package global

import (
	"apis/configs"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *configs.Config
	Log    *zap.SugaredLogger
	Db     *gorm.DB
	Engine *gin.Engine
)
