package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlc4u/health-check/internal/model"
	"github.com/rlc4u/health-check/internal/service"
	"github.com/rlc4u/health-check/pkg/logger"
	"github.com/rs/zerolog/log"
)

type HealthControllerImpl struct {
	healthSvc service.HealthSvc
}

func NewHealthController(healthSvc service.HealthSvc) HealthControllerImpl {
	return HealthControllerImpl{
		healthSvc: healthSvc,
	}
}

// swagger:route GET /health health HealthRequest
//
// Lists whether ping is okay.
//     Consumes:
//     - application/json
//     Produces:
//     - application/json
//     Schemes: http
//     Responses:
//       200: HealthResponse
func (h *HealthControllerImpl) GetHealth(ctx *gin.Context) {
	var err error
	logger.Info("Getting Query Parameters")
	query := new(model.InputQuery)
	if err = query.ParseValidate(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("True Value is %d", query.IsListening))

	data, err := h.healthSvc.HealthCheck()
	if err != nil {
		return
	}

	ctx.JSON(data.Status, data)
}
