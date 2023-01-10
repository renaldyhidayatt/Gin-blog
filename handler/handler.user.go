package handler

import (
	"fmt"
	"ginBlog/entity"
	"ginBlog/helpers"
	"ginBlog/schemas"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handlerUser struct {
	user entity.EntityUser
}

func NewHandlerUser(user entity.EntityUser) *handlerUser {
	return &handlerUser{user: user}
}

func (h *handlerUser) HandlerHello(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Hello UserRoutes", http.StatusOK, nil)
}

func (h *handlerUser) HandlerResults(ctx *gin.Context) {

	res, err := h.user.EntityResults()

	if err.Type == "error_results_01" {
		helpers.APIResponse(ctx, "User not found ", err.Code, nil)
		return
	}
	helpers.APIResponse(ctx, "User found", http.StatusOK, res)

}

func (h *handlerUser) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemasUser
	id := ctx.Param("id")
	body.ID = id

	res, err := h.user.EntityResult(&body)

	if err.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Outlet data not found for this id %s ", id), err.Code, nil)
		return

	}
	helpers.APIResponse(ctx, "USer data already to use", http.StatusOK, res)
}

func (h *handlerUser) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemasUser
	validate := validator.New()

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	if err = validate.Struct(&body); err != nil {
		helpers.APIResponse(ctx, "Invalid Validation", http.StatusBadRequest, nil)
	}

	_, error := h.user.EntityCreate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "User email already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "failed create user", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new User successfully", http.StatusCreated, nil)
}

func (h *handlerUser) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemasUser
	file, _ := ctx.FormFile("image")

	body.FirstName = ctx.PostForm("firstname")
	body.LastName = ctx.PostForm("lastname")
	body.Bio = ctx.PostForm("bio")
	body.Image = file.Filename
	body.Email = ctx.PostForm("email")
	body.Password = ctx.PostForm("password")

	err := ctx.SaveUploadedFile(file, path.Join("images/"+file.Filename))

	if err != nil {
		helpers.APIResponse(ctx, "File upload error", http.StatusBadRequest, nil)
		return
	}

	validate := validator.New()

	err = ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	if err = validate.Struct(&body); err != nil {
		helpers.APIResponse(ctx, "Invalid Validation", http.StatusBadRequest, nil)
	}

	_, error := h.user.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Failed get id", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "failed update user", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Update User successfully", http.StatusCreated, nil)
}

func (h *handlerUser) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemasUser
	id := ctx.Param("id")
	body.ID = id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	res, error := h.user.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("User data not found for this id %s ", id), error.Code, nil)
		return

	}
	helpers.APIResponse(ctx, "USer data already to use", http.StatusOK, res)
}
