package auth

import (
	"Elasticsearch/internal/pkg/constErr"
	"Elasticsearch/internal/pkg/data"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret string = "Secret Code"

func GenerateToken(name string, pass string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["username"] = name                             // Name
	atClaims["number"] = 111222333                          // RandPassword
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Time to exp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(tok string) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tok, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return constErr.NoValidToken
	}
	return nil
}

func ContainsToken(token string, baseAcc []*data.Account) (*data.Account, error) {
	for _, acc := range baseAcc {
		if acc.Token == token {
			return acc, nil
		}
	}
	return nil, constErr.TokenNotContain
}
