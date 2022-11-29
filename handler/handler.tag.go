package handler

import (
	"fmt"
	"ginBlog/entity"
	"ginBlog/helpers"
	"ginBlog/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handlerTag struct {
	tag entity.EntityTag
}

func NewHandlerTag(tag entity.EntityTag) *handlerTag {
	return &handlerTag{tag: tag}
}

func (h *handlerTag) HandlerHello(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Hello Tag Routes", http.StatusOK, nil)
}

func (h *handlerTag) HandlerResults(ctx *gin.Context) {
	res, err := h.tag.EntityResults()

	if err.Type == "error_results_01" {
		helpers.APIResponse(ctx, "User not found ", err.Code, nil)
		return
	}
	helpers.APIResponse(ctx, "Tag found", http.StatusOK, res)
}

func (h *handlerTag) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaTag
	id := ctx.Param("id")
	body.ID = id

	res, err := h.tag.EntityResult(&body)

	if err.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Outlet data not found for this id %s ", id), err.Code, nil)
		return

	}
	helpers.APIResponse(ctx, "USer data already to use", http.StatusOK, res)
}

func (h *handlerTag) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaTag
	validate := validator.New()

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	if err = validate.Struct(&body); err != nil {
		helpers.APIResponse(ctx, "Invalid Validation", http.StatusBadRequest, nil)
	}

	_, error := h.tag.EntityCreate(&body)

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

func (h *handlerTag) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaTag
	validate := validator.New()

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	if err = validate.Struct(&body); err != nil {
		helpers.APIResponse(ctx, "Invalid Validation", http.StatusBadRequest, nil)
	}

	_, error := h.tag.EntityUpdate(&body)

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

func (h *handlerTag) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaTag
	id := ctx.Param("id")
	body.ID = id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	res, error := h.tag.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("User data not found for this id %s ", id), error.Code, nil)
		return

	}
	helpers.APIResponse(ctx, "USer data already to use", http.StatusOK, res)
}
