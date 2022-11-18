package utils

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v4"
)

type AccessToken struct {
	ID    string
	Email string
}

func ExtractToken(claimsToken *jwt.Token) AccessToken {
	data := claimsToken.Claims.(jwt.MapClaims)
	parseToken := make(map[string]interface{})
	var extractToken AccessToken

	for _, v := range data {
		stringify, _ := json.Marshal(&v)
		json.Unmarshal([]byte(stringify), &parseToken)

	}

	stringify, _ := json.Marshal(&parseToken)
	json.Unmarshal([]byte(stringify), &extractToken)

	return extractToken
}
