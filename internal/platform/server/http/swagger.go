package http

import (
	_ "github.com/TrueFlowDev/Backend/docs/swagger"
	"github.com/TrueFlowDev/Backend/internal/platform/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	if cfg.App.Mode == "dev" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
