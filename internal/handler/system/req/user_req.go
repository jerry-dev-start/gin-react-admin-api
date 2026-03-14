package req

import (
	"apis/internal/model/system_model"
	"time"
)

type UserReq struct {
	Username    string     `json:"username" binding:"required"`
	Password    string     `json:"password" binding:"required"`
	Nickname    string     `json:"nickname" binding:"required"`
	PhoneNumber string     `json:"phoneNumber"`
	AvatarUrl   string     `json:"avatarUrl"`
	Birthday    *time.Time `json:"birthday"`
}

// GenDatabaseModel 生成保存到数据库的对象
func (u *UserReq) GenDatabaseModel() *system_model.User {

	return nil
}
