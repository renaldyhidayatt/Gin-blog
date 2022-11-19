package handler

import (
	"fmt"
	"ginBlog/entity"
	"ginBlog/helpers"
	"ginBlog/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerComment struct {
	comment entity.EntityComment
}

func NewHandlerComment(comment entity.EntityComment) *handlerComment {
	return &handlerComment{comment: comment}
}

func (h *handlerComment) HandlerHello(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Dota", http.StatusOK, nil)
}

func (h *handlerComment) HandlerResults(ctx *gin.Context) {
	res, err := h.comment.EntityResults()

	if err.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Comment not found ", err.Code, nil)
		return
	}
	helpers.APIResponse(ctx, "Comment found", http.StatusOK, res)
}

func (h *handlerComment) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaComment
	id := ctx.Param("id")
	body.ID = id

	res, err := h.comment.EntityResult(&body)

	if err.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Comment data not found for this id %s ", id), err.Code, nil)
		return

	}
	helpers.APIResponse(ctx, "Comment data already to use", http.StatusOK, res)
}

func (h *handlerComment) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaComment

	body.UserID = ctx.MustGet("userID").(string)

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	_, error := h.comment.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "failed create Comment", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Comment successfully", http.StatusCreated, nil)
}

func (h *handlerComment) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaComment

	body.UserID = ctx.MustGet("userID").(string)

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	_, error := h.comment.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Failed get id", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "failed update comment", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Comment successfully", http.StatusCreated, nil)
}

func (h *handlerComment) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaComment
	id := ctx.Param("id")
	body.ID = id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	res, error := h.comment.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Comment data not found for this id %s ", id), error.Code, nil)
		return

	}
	helpers.APIResponse(ctx, "Comment data already to use", http.StatusOK, res)
}
