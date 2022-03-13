package main

import (
	"github.com/gin-gonic/gin"
	"rest-api-jwt/controllers"
	"rest-api-jwt/models"
)

func main() {
	models.ConnectDatabase()
	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	r.Run(":8080")
}
