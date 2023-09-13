package router

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/API-Gopportunities/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/override/docs"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOppeningHandler)

		v1.POST("/opening", handler.CreateOppeningHandler)

		v1.DELETE("/opening", handler.DeleteOppeningHandler)

		v1.PUT("/opening", handler.UpdateOppeningHandler)

		v1.GET("/openings", handler.ListOppeningHandler)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
