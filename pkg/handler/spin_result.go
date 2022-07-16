package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/berikarg/fortune-wheel/api"
)

func (h *Handler) addSpinResult(c *gin.Context) {
	var input api.SpinResult
	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.spinResultService.Create(input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getAllSpinResults(c *gin.Context) {
	results, err := h.spinResultService.GetAll()
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, results)
}
