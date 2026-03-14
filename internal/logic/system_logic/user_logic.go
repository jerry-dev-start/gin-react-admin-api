package systemlogic

import (
	"apis/internal/data/system_data"
	"apis/internal/handler/system/req"
	"apis/internal/model/system_model"
	"errors"
	"fmt"
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
func (l *UserLogic) SaveUserInfo(r *req.UserReq) error {

	if err := l.validateUniqueness(r); err != nil {
		return err
	}

	// 生成保存到数据库实体的对象
	model, err := r.GenDatabaseModel()
	if err != nil {
		return fmt.Errorf("构建用户信息失败: %w", err)
	}

	if err := l.userData.SaveUserInfo(model); err != nil {
		return errors.New("保存用户失败")
	}
	return nil
}

// validateUniqueness 新增用户前的逻辑校验
func (l *UserLogic) validateUniqueness(r *req.UserReq) error {
	// 校验用户名
	if exist, _ := l.userData.CheckUserExistByUsername(*r.Username); exist {
		return errors.New("该用户名已存在")
	}
	// 校验手机号
	if r.PhoneNumber != nil {
		if exist, _ := l.userData.CheckUserExistByPhone(r.PhoneNumber); exist {
			return errors.New("该手机号已存在")
		}
	}
	return nil
}

func (l *UserLogic) GetUserInfoByUsername(username string) (*system_model.User, error) {
	return l.userData.GetUserInfoByUsername(username)
}
