package router

import (
	"go-cms/interface/handler/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() {
	router := gin.Default()
	authGroup := router.Group("/admin")

	router.GET("/")

	router.GET("/login")
	router.POST("/login")
	router.GET("/register")

	authGroup.Use(middleware.IsLoggedIn())
	{
		authGroup.GET("/dashbord")
	}

	router.Run()
}
