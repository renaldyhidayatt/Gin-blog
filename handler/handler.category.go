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

type handlerCategory struct {
	category entity.EntityCategory
}

func NewHandlerCategory(category entity.EntityCategory) *handlerCategory {
	return &handlerCategory{category: category}
}

func (h *handlerCategory) HandlerHello(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Hello Category Routes", http.StatusOK, nil)
}

func (h *handlerCategory) HandlerResults(ctx *gin.Context) {

	res, err := h.category.EntityResults()

	if err.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Category not found ", err.Code, nil)
		return
	}
	helpers.APIResponse(ctx, "Category found", http.StatusOK, res)

}

func (h *handlerCategory) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaCategories
	id := ctx.Param("id")
	body.ID = id

	res, err := h.category.EntityResult(&body)

	if err.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Category data not found for this id %s ", id), err.Code, nil)
		return

	}
	helpers.APIResponse(ctx, "Category data already to use", http.StatusOK, res)
}

func (h *handlerCategory) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaCategories
	validate := validator.New()

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	if err = validate.Struct(&body); err != nil {
		helpers.APIResponse(ctx, "Invalid Validation", http.StatusBadRequest, nil)
	}

	_, error := h.category.EntityCreate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Category name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "failed create category", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Category successfully", http.StatusCreated, nil)
}

func (h *handlerCategory) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaCategories
	validate := validator.New()

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	if err = validate.Struct(&body); err != nil {
		helpers.APIResponse(ctx, "Invalid Validation", http.StatusBadRequest, nil)
	}

	_, error := h.category.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Failed get id", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "failed update category", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Category User successfully", http.StatusCreated, nil)
}

func (h *handlerCategory) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaCategories
	id := ctx.Param("id")
	body.ID = id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	res, error := h.category.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Categpry data not found for this id %s ", id), error.Code, nil)
		return

	}
	helpers.APIResponse(ctx, "Category data already to use", http.StatusOK, res)
}
