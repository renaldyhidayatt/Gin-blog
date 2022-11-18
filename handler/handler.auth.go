package handler

import (
	"ginBlog/entity"
	"ginBlog/helpers"
	"ginBlog/schemas"
	"ginBlog/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type handlerAuth struct {
	auth entity.EntityAuth
}

func NewHandlerAuth(auth entity.EntityAuth) *handlerAuth {
	return &handlerAuth{auth: auth}
}

func (h *handlerAuth) HandlerHello(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Dota", http.StatusOK, nil)
}

func (h *handlerAuth) HandlerRegister(ctx *gin.Context) {
	var body schemas.SchemasUser

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	_, error := h.auth.EntityRegister(&body)

	if error.Type == "error_register_01" {
		helpers.APIResponse(ctx, "Email already taken", error.Code, nil)
		return
	}

	if error.Type == "error_register_02" {
		helpers.APIResponse(ctx, "Register new user account failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Register new user account success", http.StatusOK, nil)

}

func (h *handlerAuth) HandlerLogin(ctx *gin.Context) {
	var body schemas.SchemasUser
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	res, error := h.auth.EntityLogin(&body)

	if error.Type == "error_login_01" {
		helpers.APIResponse(ctx, "User account is not never registered", error.Code, nil)
		return
	}

	if error.Type == "error_login_02" {
		helpers.APIResponse(ctx, "Email or Password is wrong", error.Code, nil)
		return
	}

	accessToken, errorJwt := utils.GenerateToken(&schemas.JWtMetaRequest{
		Data:      gin.H{"id": res.ID, "email": res.Email},
		SecretKey: utils.GodotEnv("JWT_SECRET_KEY"),
		Options:   schemas.JwtMetaOptions{Audience: "majoo", ExpiredAt: 1},
	})

	expiredAt := time.Now().Add(time.Duration(time.Minute) * (24 * 60) * 1).Local()

	if errorJwt != nil {
		helpers.APIResponse(ctx, "Generate access token failed", http.StatusBadRequest, nil)
		return
	}

	helpers.APIResponse(ctx, "Login successfully", http.StatusOK, gin.H{"accessToken": accessToken, "expiredAt": expiredAt})
}
