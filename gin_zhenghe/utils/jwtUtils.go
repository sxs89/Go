package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	UserId               int64                        `json:"user_id,omitempty"`
	UserName             string                       `json:"user_name"`
	jwt.RegisteredClaims `:"jwt_._registered_claims"` // 内嵌标准说明
}

/**
生成token
*/
func GenerateJWT(userId int64, username string) (string, error) {
	claims := CustomClaims{
		userId,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 设置过期时间为24小时
			Issuer:    "sxs",                                              // 设置签发者
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 设置秘钥进行签名
	secret := []byte("sxs")
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

/**
验证Token
*/
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析并验证JWT
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证秘钥
		secret := []byte("sxs")
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid JWT")
	}

	return claims, nil
}

// 更新token时间
func UpdateJwtExpiration(tokenString string, duration time.Duration) (string, error) {
	// 解析jwt
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &CustomClaims{})
	if err != nil {
		return "", err
	}

	// 获取原始声明信息
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return "", fmt.Errorf("Invalid claims in JWT")
	}

	// 更新过期时间
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(duration))

	// 生成新的JWT
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte("sxs")
	newTokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return newTokenString, nil

}
