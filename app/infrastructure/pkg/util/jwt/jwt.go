package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"vending/app/domain/obj"
)

var config *Config

type Config struct {
	JwtSecret     string `yaml:"jwtSecret"`
	JwtExpireTime int    `yaml:"jwtExpireTime"`
	Issuer        string `yaml:"issuer"`
	Secret        []byte `yaml:"secret"`
	JwtAuthKey    string `yaml:"jwtAuthKey"`
}

func New(c *Config) {
	config = c
}

func GenerateToken(name string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.JwtExpireTime) * time.Second)

	claims := obj.Claims{
		JwtToken: obj.JwtToken{Username: name},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//  该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(config.Secret)
	return token, err
}

func ParseToken(token string) (*obj.Claims, error) {
	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &obj.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.Secret, nil
	})

	if tokenClaims != nil {
		// 验证基于时间的声明exp, iat, nbf，注意如果没有任何声明在令牌中
		// 仍然会被认为是有效的。并且对于时区偏差没有计算方法
		if claims, ok := tokenClaims.Claims.(*obj.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
