package helpers

import (
	"ginBlog/entity"
	"ginBlog/schemas"
	"math"

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

func APIResPagination(ctx *gin.Context, message string, status int, data interface{}, page int, size int, h entity.EntityArticle) {

	count, _ := h.EntityCount()

	var totalPages float64
	if float64(count) < float64(size) {
		totalPages = 1
	} else {
		totalPages = math.Ceil(float64(count) / float64(size))
	}

	pagination := schemas.Pagination{
		Page:       page,
		Size:       size,
		Count:      count,
		TotalPages: totalPages,
	}

	ctx.JSON(status, gin.H{
		"message":    message,
		"data":       data,
		"pagination": pagination,
	})
}
