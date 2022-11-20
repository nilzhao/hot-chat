package utils

import (
	"hot-chat/global"
	"hot-chat/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type BaseClaims struct {
	model.User
}

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

func createClaims(baseClaims BaseClaims) CustomClaims {
	claims := CustomClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Duration(-1000) * time.Second)),                    // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.CONFIG.Jwt.Period) * time.Second)), // 过期时间 7天  配置文件
			Issuer:    global.CONFIG.Jwt.Issuer,                                                                  // 签名的发行者
		},
	}
	return claims
}

func CreateToken(baseClaims BaseClaims) (string, int64) {
	claims := createClaims(baseClaims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(global.CONFIG.Jwt.Salt)
	if err != nil {
		global.Logger.Error(err)
		return "", 0
	}
	return tokenString, claims.ExpiresAt.Unix()
}

func ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return global.CONFIG.Jwt.Salt, nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*CustomClaims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
