package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang-telegram-bot/api"
	docs "golang-telegram-bot/docs"
)

// @title Album API
// @description Test Album API.
// @schemes http https
// @BasePath /api/v1
func main() {
	fmt.Println("Server starts  at localhost:8083 ... ")
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	apiGroup := router.Group("/api")
	{
		v1 := apiGroup.Group("/v1")
		{
			v1.GET("/albums", api.GetAlbums)
			v1.GET("/albums/:id", api.GetAlbumByID)
			v1.POST("/albums", api.PostAlbums)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8083")
}
