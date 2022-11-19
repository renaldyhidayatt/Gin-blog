package handler

import (
	"fmt"
	"ginBlog/entity"
	"ginBlog/helpers"
	"ginBlog/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerArticle struct {
	article entity.EntityArticle
}

func NewHandlerArticle(article entity.EntityArticle) *handlerArticle {
	return &handlerArticle{article: article}
}

func (h *handlerArticle) HandlerHello(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Dota", http.StatusOK, nil)
}

func (h *handlerArticle) HandlerResults(ctx *gin.Context) {
	res, err := h.article.EntityResults()

	if err.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Comment not found ", err.Code, nil)
		return
	}
	helpers.APIResponse(ctx, "Article found", http.StatusOK, res)
}

func (h *handlerArticle) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaArticle
	id := ctx.Param("id")
	body.ID = id

	res, err := h.article.EntityResult(&body)

	if err.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Article data not found for this id %s ", id), err.Code, nil)
		return

	}
	helpers.APIResponse(ctx, "Comment data already to use", http.StatusOK, res)
}

func (h *handlerArticle) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaArticle

	body.UserID = ctx.MustGet("userID").(string)

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	_, error := h.article.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "failed Article Comment", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Article successfully", http.StatusCreated, nil)
}

func (h *handlerArticle) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaArticle

	body.UserID = ctx.MustGet("userID").(string)

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	_, error := h.article.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Failed get id", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "failed update comment", error.Code, nil)
		return
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

	res, error := h.article.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Article data not found for this id %s ", id), error.Code, nil)
		return

	}
	helpers.APIResponse(ctx, "Article data already to use", http.StatusOK, res)
}
