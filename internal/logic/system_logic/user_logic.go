package systemlogic

import (
	"apis/internal/data/system_data"
	"apis/internal/handler/system/req"
	"errors"
)

type UserLogic struct {
	userData *system_data.UserData
}

func NewUserLogic() *UserLogic {
	return &UserLogic{
		userData: system_data.NewUserData(),
	}
}

// SaveUserInfo 保存用户信息
func (l UserLogic) SaveUserInfo(r *req.UserReq) error {
	exist, err := l.userData.CheckUserExistByUsername(r.Username)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("该用户名已存在")
	}

	// 生成保存到数据库实体的对象

	return exist
}
