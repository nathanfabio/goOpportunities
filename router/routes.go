package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/goOpportunities/handler"
	docs "github.com/nathanfabio/goOpportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize handler
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningHandler)
	}

	// Swagger
	router.GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
}
