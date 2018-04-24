package server

import (
	"Go-GoSAFE.converter/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/import/railml", controllers.ImportRailml)
		v1.POST("/export/railml", controllers.ExportRailml)
	}

	return router // listen and serve on 0.0.0.0:8080
}
