package http

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(dataSVCHandler *DataSVCHandler) *gin.Engine {
	r := gin.Default()
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.Use(middleware.CORS())
	dataSVC := r.Group("/data-service")
	{
		dataSVC.GET("coin_details", dataSVCHandler.GetCoinDetails)
	}

	return r
}
