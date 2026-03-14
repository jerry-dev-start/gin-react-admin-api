package system_data

import (
	"apis/global"
	"apis/internal/model/system_model"
	"errors"

	"gorm.io/gorm"
)

type UserData struct {
}

func NewUserData() *UserData {
	return &UserData{}
}

func (d *UserData) CheckUserExistByUsername(username string) (bool, error) {
	var userModel system_model.User
	if err := global.Db.Where("username = ?", username).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (d *UserData) SaveUserInfo(model *system_model.User) error {
	if err := global.Db.Create(model).Error; err != nil {
		return err
	}
	return nil
}

func (d *UserData) CheckUserExistByPhone(number *string) (bool, error) {
	var userModel system_model.User
	if err := global.Db.Where("phone_number = ?", number).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (d *UserData) GetUserInfoByUsername(username string) (*system_model.User, error) {
	var userInfo system_model.User
	err := global.Db.Where("username = ?", username).First(&userInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &userInfo, err
}
