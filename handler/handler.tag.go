package handler

import (
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

	if err != nil {
		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusBadRequest, nil)
		return
	}
	helpers.APIResponse(ctx, "Tag found", http.StatusOK, res)
}

func (h *handlerTag) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaTag
	id := ctx.Param("id")
	body.ID = id

	res, err := h.tag.EntityResult(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusBadRequest, nil)
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

	_, err = h.tag.EntityCreate(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusBadRequest, nil)

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

	_, err = h.tag.EntityUpdate(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusBadRequest, nil)
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

	res, err := h.tag.EntityDelete(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Error: %s"+err.Error(), http.StatusBadRequest, nil)
		return

	}
	helpers.APIResponse(ctx, "USer uccessfully delete", http.StatusOK, res)
}
