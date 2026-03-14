package pkg

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID               uint64 `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 包含过期时间、发行人等标准字段
}

// 签名密钥（生产环境应从配置文件或环境变量读取）
var jwtSecret = []byte("your_very_secret_key_666")

// GenerateToken 生成 JWT
func GenerateToken(userID uint64, username string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置 24 小时过期
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-app",
		},
	}

	// 使用 HS256 算法创建 Token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥进行签名并返回字符串形式的 Token
	return token.SignedString(jwtSecret)
}

// ParseToken 解析并验证 JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 校验解密后的内容并返回
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
