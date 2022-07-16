package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/berikarg/fortune-wheel/api"
	"github.com/berikarg/fortune-wheel/pkg/service"
)

type Handler struct {
	spinResultService service.SpinResult
	Logger            *zap.Logger
	Wheel             api.Wheel
}

func NewHandler(services service.SpinResult, logger *zap.Logger, initWheel api.Wheel) *Handler {
	return &Handler{
		spinResultService: services,
		Logger:            logger,
		Wheel:             initWheel,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	apiGroup := router.Group("/api")
	{
		wheel := apiGroup.Group("/wheel")
		{
			wheel.GET("/", h.getWheel)
		}

		spinResults := apiGroup.Group("/results")
		{
			spinResults.POST("/", h.addSpinResult)
		}
	}

	admin := router.Group("/admin")
	{
		wheel := admin.Group("/wheel")
		{
			wheel.POST("/", h.setWheel)
		}
		spinResults := admin.Group("/results")
		{
			spinResults.GET("/", h.getAllSpinResults)
		}
	}

	return router
}
