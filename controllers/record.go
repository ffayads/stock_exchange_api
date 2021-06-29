package controllers

import (
	"net/http"

	"github.com/api/stock_exchange_api/helpers"
	"github.com/api/stock_exchange_api/httpmodels"
	"github.com/api/stock_exchange_api/models"
	"github.com/gin-gonic/gin"
)

func GetRecords(c *gin.Context) {
	user := c.Value("user").(*models.Users)
	filter := &httpmodels.Filter{}
	if err := c.Bind(filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Datos Incompletos", "data": gin.H{}})
		return
	}
	record := helpers.GetRecords(filter, user)
	if record == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "datos no encontrados", "data": gin.H{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "",
		"data":    record,
	})
	return
}

func GetRecordDetail(c *gin.Context) {
	user := c.Value("user").(*models.Users)
	filter := &httpmodels.GetRecordDetail{}
	if err := c.Bind(filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Datos Incompletos", "data": gin.H{}})
		return
	}
	record := helpers.GetRecord(filter, user)
	if record == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "datos no encontrados", "data": gin.H{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "",
		"data":    record,
	})
	return
}
