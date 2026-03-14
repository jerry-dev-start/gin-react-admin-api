package req

import (
	"apis/internal/bootstrap/snowflake"
	"apis/internal/model/system_model"
	"apis/pkg"
	"time"
)

type UserReq struct {
	Username    *string    `json:"username" binding:"required"`
	Password    *string    `json:"password" binding:"required"`
	Nickname    *string    `json:"nickname" binding:"required"`
	PhoneNumber *string    `json:"phoneNumber"`
	AvatarUrl   *string    `json:"avatarUrl"`
	Birthday    *time.Time `json:"birthday"`
	Email       *string    `json:"email"`
}

// GenDatabaseModel 生成保存到数据库的对象
func (u *UserReq) GenDatabaseModel() (*system_model.User, error) {

	ur := &system_model.User{
		Username:     u.Username,
		Nickname:     u.Nickname,
		PhoneNumber:  u.PhoneNumber,
		AvatarURL:    u.AvatarUrl,
		PasswordHash: u.Password,
		Birthday:     u.Birthday,
		ID:           snowflake.GenID(),
		Email:        u.Email,
	}
	if ur.PasswordHash != nil {
		hashed, err := pkg.HashPassword(*ur.PasswordHash)
		if err != nil {
			return nil, err
		}
		ur.PasswordHash = &hashed
	}
	return ur, nil
}
