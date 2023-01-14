package handler

import (
	"ginBlog/entity"
	"ginBlog/helpers"
	"ginBlog/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handlerComment struct {
	comment entity.EntityComment
}

func NewHandlerComment(comment entity.EntityComment) *handlerComment {
	return &handlerComment{comment: comment}
}

func (h *handlerComment) HandlerHello(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Hello Comment Routes", http.StatusOK, nil)
}

func (h *handlerComment) HandlerResults(ctx *gin.Context) {
	res, err := h.comment.EntityResults()

	if err != nil {
		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusBadRequest, nil)
		return
	}
	helpers.APIResponse(ctx, "Comment found", http.StatusOK, res)
}

func (h *handlerComment) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaComment
	id := ctx.Param("id")
	body.ID = id

	res, err := h.comment.EntityResult(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusBadRequest, nil)
		return

	}
	helpers.APIResponse(ctx, "Comment data already to use", http.StatusOK, res)
}

func (h *handlerComment) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaComment
	validate := validator.New()

	body.UserID = ctx.MustGet("userID").(string)

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	if err = validate.Struct(&body); err != nil {
		helpers.APIResponse(ctx, "Invalid Validation", http.StatusBadRequest, nil)
	}

	_, err = h.comment.EntityCreate(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusBadGateway, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Comment successfully", http.StatusCreated, nil)
}

func (h *handlerComment) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaComment
	validate := validator.New()

	body.UserID = ctx.MustGet("userID").(string)

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json body failed", http.StatusBadRequest, nil)
	}

	if err = validate.Struct(&body); err != nil {
		helpers.APIResponse(ctx, "Invalid Validation", http.StatusBadRequest, nil)
	}

	_, err = h.comment.EntityUpdate(&body)

	if err != nil {

		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusBadGateway, nil)
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

	res, err := h.comment.EntityDelete(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusBadRequest, nil)
		return

	}
	helpers.APIResponse(ctx, "Comment successfully delete", http.StatusOK, res)
}
