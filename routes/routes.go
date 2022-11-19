package routes

import (
	"ginBlog/handler"
	"ginBlog/repository"
	"ginBlog/service"

	"ginBlog/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRoute(db *gorm.DB, router *gin.Engine) {
	repositoryAuth := repository.NewRepositoryAuth(db)
	serviceAuth := service.NewServiceAuth(repositoryAuth)
	handlerAuth := handler.NewHandlerAuth(serviceAuth)

	repositoryUser := repository.NewRepositoryUser(db)
	serviceUser := service.NewServiceUser(repositoryUser)
	handlerUser := handler.NewHandlerUser(serviceUser)

	route := router.Group("/api/auth")
	routeUser := router.Group("/api/user")

	route.POST("/ping", handlerAuth.HandlerHello)
	route.POST("/register", handlerAuth.HandlerRegister)
	route.POST("/login", handlerAuth.HandlerLogin)

	routeUser.Use(middleware.AuthToken())

	routeUser.GET("/", handlerUser.HandlerResults)
	routeUser.POST("/create", handlerUser.HandlerCreate)
	routeUser.GET("/:id", handlerUser.HandlerResult)
	routeUser.PUT("/:id", handlerUser.HandlerUpdate)
	routeUser.DELETE("/:id", handlerUser.HandlerResults)
}
