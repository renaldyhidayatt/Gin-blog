package utils

import (
	"ginBlog/schemas"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

func GenerateToken(configs *schemas.JWtMetaRequest) (string, error) {
	expiredAt := time.Now().Add(time.Duration(time.Minute) * configs.Options.ExpiredAt).Unix()

	claims := jwt.MapClaims{}
	claims["jwt"] = configs.Data
	claims["exp"] = (24 * 60) * expiredAt
	claims["audience"] = configs.Options.Audience
	claims["authorization"] = true

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(configs.SecretKey))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

func VerifyToken(accessToken, SecretPublicKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretPublicKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
	}

	return token, nil
}
