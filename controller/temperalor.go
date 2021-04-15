package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kashimashino/hackweek-group6/service"
	"net/http"
)

func Temperature(c *gin.Context) {
	var tpa service.TempCalcService

	tpa.Username = c.Param("username")

	if url, err := tpa.TempCalc(); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, url)
	}
}
