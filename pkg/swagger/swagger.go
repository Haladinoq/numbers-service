package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.palo-it.net/palo/numbers-service/pkg/config"
)

func InitSwagger(router *gin.Engine, config *config.Config) {

	swaggerGroup := router.Group("/swagger")
	swaggerGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
