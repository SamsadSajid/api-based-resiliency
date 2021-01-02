package router

import (
	"github.com/gin-gonic/gin"
	"github.com/resilience-poc/pay-bill-resilience-service/controller"
)

func New(router *gin.Engine) *gin.Engine {
	base := router.Group("/api/")

	base.POST("/update-biller-health", controller.HealthController)

	return router
}
