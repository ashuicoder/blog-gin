package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(userId uint, username string) (string, error) {
	// 创建一个我们自己声明的数据
	claims := CustomClaims{
		userId,
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 定义过期时间
			// 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString([]byte("secret-key"))
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte("secret-key"), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok {
		return nil, err
	}

	return claims, nil
}
