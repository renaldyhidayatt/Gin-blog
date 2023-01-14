package handler

import (
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

	if err != nil {
		helpers.APIResponse(ctx, "Category not found ", http.StatusNotFound, nil)
		return
	}
	helpers.APIResponse(ctx, "Category found", http.StatusOK, res)

}

func (h *handlerCategory) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaCategories
	id := ctx.Param("id")
	body.ID = id

	res, err := h.category.EntityResult(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusNotFound, nil)
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

	_, err = h.category.EntityCreate(&body)

	if err != nil {

		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusBadRequest, nil)
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

	_, err = h.category.EntityUpdate(&body)

	if err != nil {
		helpers.APIResponse(ctx, err.Error(), http.StatusBadRequest, nil)

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

	res, err := h.category.EntityDelete(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusBadRequest, nil)
		return

	}
	helpers.APIResponse(ctx, "Category successfully delete", http.StatusOK, res)
}
