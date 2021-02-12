package router

import (
	"github.com/george124816/api-vuttr/internal/api/controllers"
	"github.com/gin-gonic/gin"

	_ "github.com/george124816/api-vuttr/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"

)

func Setup() *gin.Engine{
	router := gin.New()

	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/tools", controllers.GetTools)
		v1.POST("/tools", controllers.InsertTool)
		v1.DELETE("/tools/:id", controllers.DeleteTool)
	}

	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, ""))

	return router
}