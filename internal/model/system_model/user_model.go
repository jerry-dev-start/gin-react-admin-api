package system_model

import "time"

type User struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	SsoUserId string `gorm:"type:varchar(100);unique;column:sso_user_id;" json:"user_uuid"`

	// 登录核心信息
	Username     *string `gorm:"type:varchar(50);unique;not null;column:username" json:"username"`
	Email        *string `gorm:"type:varchar(100);unique;column:email" json:"email"`
	PhoneNumber  *string `gorm:"type:varchar(20);unique;column:phone_number" json:"phone_number"`
	PasswordHash *string `gorm:"type:varchar(255);not null;column:password_hash" json:"-"` // 密码哈希通常在 JSON 中隐藏

	// 个人档案
	Nickname  *string    `gorm:"type:varchar(50);column:nickname" json:"nickname"`
	AvatarURL *string    `gorm:"type:varchar(255);column:avatar_url" json:"avatar_url"`
	Gender    int8       `gorm:"type:tinyint;default:0;column:gender" json:"gender"` // 0:未知, 1:男, 2:女
	Birthday  *time.Time `gorm:"type:date;column:birthday" json:"birthday"`
	Bio       string     `gorm:"type:text;column:bio" json:"bio"`

	// 账号状态
	Status    int8 `gorm:"type:tinyint;default:1;column:status" json:"status"`
	IsDeleted int8 `gorm:"type:tinyint(1);default:0;column:is_deleted" json:"is_deleted"`

	// 时间戳 (对应 MySQL 的 TIMESTAMP)
	LastLoginAt *time.Time `gorm:"column:last_login_at" json:"last_login_at"`
	LastLoginIP *string    `gorm:"type:varchar(45);column:last_login_ip" json:"last_login_ip"`
	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName 指定 GORM 使用的表名
func (User) TableName() string {
	return "users"
}
