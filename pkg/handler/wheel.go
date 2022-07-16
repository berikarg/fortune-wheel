package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/berikarg/fortune-wheel/api"
)

func (h *Handler) getWheel(c *gin.Context) {
	c.JSON(http.StatusOK, h.Wheel)
}

func (h *Handler) setWheel(c *gin.Context) {
	var input api.Wheel
	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	h.Wheel = input
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
