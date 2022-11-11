package router

import (
	"bookmanager/controller"
	"bookmanager/middleware"
	"github.com/gin-gonic/gin"
)

func SetupApiRouters(r *gin.Engine) {
	r.POST("/api/register", controller.RegisterHandler)
	r.POST("/api/login", controller.LoginHandler)
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())
	v1.POST("/book", controller.CreateBookHandler)
	v1.GET("/books", controller.GetBookListHandler)
	v1.GET("/book/:id", controller.GetBookDetailHandler)
	v1.PUT("/book", controller.UpdateBookHandler)
	v1.DELETE("/book/:id", controller.DeleteBookHandler)
}
