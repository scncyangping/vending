package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"vending/data/common/config"
)

type Claims struct {
	Username string `json:"username"`
	UserId   string `json:"userid"`
	jwt.StandardClaims
}

func GenerateToken(username, id string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.Base.Jwt.JwtExpireTime) * time.Second)

	claims := Claims{
		username,
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.Base.Jwt.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//  该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(config.Base.Jwt.Secret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.Base.Jwt.Secret, nil
	})

	if tokenClaims != nil {
		// 验证基于时间的声明exp, iat, nbf，注意如果没有任何声明在令牌中
		// 仍然会被认为是有效的。并且对于时区偏差没有计算方法
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
