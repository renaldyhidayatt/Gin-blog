package helpers

import (
	"ginBlog/schemas"

	"github.com/gin-gonic/gin"
)

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Data interface{}) {

	jsonResponse := schemas.SchemaResponses{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func ErrorResponse(ctx *gin.Context, StatusCode int, Error string) {
	err := schemas.SchemaErrorResponse{
		StatusCode: StatusCode,
		Message:    Error,
	}

	ctx.AbortWithStatusJSON(err.StatusCode, err)
}
