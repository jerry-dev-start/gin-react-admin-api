package pkg

import "golang.org/x/crypto/bcrypt"

// HashPassword 对明文密码进行 Bcrypt 加密
func HashPassword(password string) (string, error) {
	// GenerateFromPassword 的第二个参数是 Cost，范围 4-31，默认 10 即可
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPasswordHash 验证明文密码与哈希值是否匹配
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
