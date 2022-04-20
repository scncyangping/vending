package obj

import "github.com/dgrijalva/jwt-go"

type JwtToken struct {
	Username string `json:"username"`
}

type Claims struct {
	JwtToken
	jwt.StandardClaims
}

type Permission struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}
