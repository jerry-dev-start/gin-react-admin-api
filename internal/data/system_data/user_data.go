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
