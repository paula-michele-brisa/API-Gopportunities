package router

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/API-Gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.CreateOppeningHandler)

		v1.POST("/opening", handler.ShowOppeningHandler)

		v1.DELETE("/opening", handler.DeleteOppeningHandler)

		v1.PUT("/opening", handler.UpdateOppeningHandler)

		v1.GET("/openings", handler.ListOppeningHandler)

	}

}
