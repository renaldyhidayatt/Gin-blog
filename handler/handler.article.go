package handler

import (
	"ginBlog/entity"
	"ginBlog/helpers"
	"ginBlog/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handlerArticle struct {
	article entity.EntityArticle
}

func NewHandlerArticle(article entity.EntityArticle) *handlerArticle {
	return &handlerArticle{article: article}
}

func (h *handlerArticle) HandlerHello(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Hello Article", http.StatusOK, nil)
}

func (h *handlerArticle) HandlerResults(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	res, err := h.article.EntityResults(page, size)

	if err != nil {
		helpers.APIResponse(ctx, err.Error(), http.StatusNotFound, nil)

	}

	helpers.APIResPagination(ctx, "berhasil mendapatkan data", http.StatusOK, res, page, size, h.article)
}

func (h *handlerArticle) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaArticle
	id := ctx.Param("id")
	body.ID = id

	res, err := h.article.EntityResult(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Error %s"+err.Error(), http.StatusNotFound, nil)
		return

	}

	helpers.APIResponse(ctx, "Comment data already to use", http.StatusOK, res)
}

func (h *handlerArticle) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaArticle
	validate := validator.New()

	body.UserID = ctx.MustGet("userID").(string)

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	if err = validate.Struct(&body); err != nil {
		helpers.APIResponse(ctx, "Invalid Validation", http.StatusBadRequest, nil)
	}

	_, err = h.article.EntityCreate(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Error %s"+err.Error(), http.StatusBadRequest, nil)
		return

	}

	helpers.APIResponse(ctx, "Create new Article successfully", http.StatusCreated, nil)
}

func (h *handlerArticle) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaArticle
	validate := validator.New()

	body.UserID = ctx.MustGet("userID").(string)

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	if err = validate.Struct(&body); err != nil {
		helpers.APIResponse(ctx, "Invalid Validation", http.StatusBadRequest, nil)
	}

	_, err = h.article.EntityUpdate(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Error %s"+err.Error(), http.StatusNotFound, nil)
	}

	helpers.APIResponse(ctx, "Article successfully", http.StatusCreated, nil)
}

func (h *handlerArticle) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaArticle
	id := ctx.Param("id")
	body.ID = id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	res, err := h.article.EntityDelete(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusNotFound, nil)
	}
	helpers.APIResponse(ctx, "Article successfully delete", http.StatusOK, res)
}
