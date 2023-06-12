package jwta

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserInfo struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type JWTClaims struct {
	User UserInfo
	jwt.RegisteredClaims
}

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("88888888")

func GenerateToken(userInfo UserInfo) (string, error) {
	expirationTime := time.Now().Add(TokenExpireDuration) // 两个小时有效期
	claims := &JWTClaims{
		User: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	// 生成Token，指定签名算法和claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名
	if tokenString, err := token.SignedString(MySecret); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func ParseToken(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// 检查过期时间
	if time.Now().After(claims.ExpiresAt.Time) {
		return nil, errors.New("token is expired")
	}

	return claims, nil
}
