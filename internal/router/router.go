package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlc4u/health-check/internal/controller"
	"github.com/rlc4u/health-check/internal/model"
	"github.com/rlc4u/health-check/internal/service"
	"github.com/rlc4u/health-check/pkg/client"
	"github.com/rlc4u/health-check/pkg/config"
	"github.com/rlc4u/health-check/pkg/logger"
)

func SetupRouter() *gin.Engine {
	conf := config.InitReader()
	logger.Setup("health-check-service", conf.LOGGING_LEVEL)
	logger.Debug("Will not appear")
	logger.Info("Checkpoint 1")
	r := gin.Default()

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, model.Response{
			Status: http.StatusNotFound,
			Data:   "No Such Service",
		})
	})

	httpClient := client.NewClient()
	healthSvc := service.NewHealthSvc(httpClient)
	healthController := controller.NewHealthController(healthSvc)

	r.GET("health", healthController.GetHealth)

	return r
}
