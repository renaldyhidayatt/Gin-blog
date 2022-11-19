package middleware

import (
	"ginBlog/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthToken() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		bearer := ctx.GetHeader("Authorization")

		if bearer == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
			return
		}

		token := strings.Split(bearer, " ")

		if len(token) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Access Token is required"})
			return
		}

		decodeToken, err := utils.VerifyToken(strings.TrimSpace(token[1]), utils.GodotEnv("JWT_SECRET_KEY"))

		if err != nil {
			defer logrus.Error(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Access Token Expired"})
		}

		accessToken := utils.ExtractToken(decodeToken)

		ctx.Set("user", accessToken)
		ctx.Set("userID", accessToken.ID)
		ctx.Next()

	})
}
