package main

import (
	"github.com/gin-gonic/gin"
	"rest-api-jwt/controllers"
	"rest-api-jwt/middlewares"
	"rest-api-jwt/models"
)

func main() {
	models.ConnectDatabase()
	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	r.Run(":8080")
}
