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

	repositoryCategory := repository.NewRepositoryCategory(db)
	serviceCategory := service.NewServiceCategory(repositoryCategory)
	handlerCategory := handler.NewHandlerCategory(serviceCategory)

	repositoryTag := repository.NewRepositoryTag(db)
	serviceTag := service.NewServiceTag(repositoryTag)
	handlerTag := handler.NewHandlerTag(serviceTag)

	repositoryArticle := repository.NewRepositoryArticle(db)
	serviceArticle := service.NewServiceArticle(repositoryArticle)
	handlerArticle := handler.NewHandlerArticle(serviceArticle)

	repositoryComment := repository.NewRepositoryComment(db)
	serviceComment := service.NewServiceComment(repositoryComment)
	handlerComment := handler.NewHandlerComment(serviceComment)

	route := router.Group("/api/auth")
	routeUser := router.Group("/api/user")
	routeCategory := router.Group("/api/category")
	routeTag := router.Group("/api/tag")
	routeArticle := router.Group("/api/article")
	routerComment := router.Group("/api/comment")

	route.POST("/ping", handlerAuth.HandlerHello)
	route.POST("/register", handlerAuth.HandlerRegister)
	route.POST("/login", handlerAuth.HandlerLogin)

	routeUser.Use(middleware.AuthToken())
	routeCategory.Use(middleware.AuthToken())
	routeTag.Use(middleware.AuthToken())
	routeArticle.Use(middleware.AuthToken())
	routerComment.Use(middleware.AuthToken())

	routeUser.GET("/", handlerUser.HandlerResults)
	routeUser.POST("/create", handlerUser.HandlerCreate)
	routeUser.GET("/:id", handlerUser.HandlerResult)
	routeUser.PUT("/:id", handlerUser.HandlerUpdate)
	routeUser.DELETE("/:id", handlerUser.HandlerResults)

	routeCategory.GET("/", handlerCategory.HandlerResults)
	routeCategory.POST("/create", handlerCategory.HandlerCreate)
	routeCategory.GET("/:id", handlerCategory.HandlerResult)
	routeCategory.PUT("/:id", handlerCategory.HandlerUpdate)
	routeCategory.DELETE("/:id", handlerCategory.HandlerResults)

	routeTag.GET("/", handlerTag.HandlerResults)
	routeTag.POST("/create", handlerTag.HandlerCreate)
	routeTag.GET("/:id", handlerTag.HandlerResult)
	routeTag.PUT("/:id", handlerTag.HandlerUpdate)
	routeTag.DELETE("/:id", handlerTag.HandlerResults)

	routeArticle.GET("/", handlerArticle.HandlerResults)
	routeArticle.POST("/create", handlerArticle.HandlerCreate)
	routeArticle.GET("/:id", handlerArticle.HandlerResult)
	routeArticle.PUT("/:id", handlerArticle.HandlerUpdate)
	routeArticle.DELETE("/:id", handlerArticle.HandlerResults)

	routerComment.GET("/", handlerArticle.HandlerResults)
	routerComment.POST("/create", handlerArticle.HandlerCreate)
	routerComment.GET("/:id", handlerArticle.HandlerResult)
	routerComment.PUT("/:id", handlerArticle.HandlerUpdate)
	routerComment.DELETE("/:id", handlerArticle.HandlerResults)
}
