package utils

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtParser struct {
	Tokens map[string]interface{}
}

type Claims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

func NewJwtParser() *JwtParser {
	return &JwtParser{}
}

func (j *JwtParser) GetTokens(optr string) map[string]interface{} {
	tokens := map[string]interface{}{
		"optr": "openapi_secret",
	}
	j.Tokens = tokens
	return tokens
}

func (j *JwtParser) GenToken(optr, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		optr,
		jwt.StandardClaims{},
	})
	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

func (j *JwtParser) IsValid(tokenString string) (bool, *Claims) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			claims := token.Claims.(*Claims)
			secret := j.Tokens[claims.Account].(string)
			if secret == "" {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		},
	)

	if err != nil {
		return false, nil
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return true, claims
	} else {
		return false, nil
	}
}
