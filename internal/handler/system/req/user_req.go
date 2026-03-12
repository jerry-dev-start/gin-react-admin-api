package req

import "time"

type UserReq struct {
	Username    string     `json:"username" binding:"required"`
	Password    string     `json:"password" binding:"required"`
	Nickname    string     `json:"nickname" binding:"required"`
	PhoneNumber string     `json:"phoneNumber"`
	AvatarUrl   string     `json:"avatarUrl"`
	Birthday    *time.Time `json:"birthday"`
}
