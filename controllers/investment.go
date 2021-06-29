package controllers

import (
	"net/http"

	"github.com/api/stock_exchange_api/helpers"
	"github.com/api/stock_exchange_api/httpmodels"
	"github.com/api/stock_exchange_api/models"
	"github.com/api/stock_exchange_api/repository"
	"github.com/gin-gonic/gin"
)

func GetInvestments(c *gin.Context) {
	u := &httpmodels.Filter{}
	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Datos Incompletos", "data": gin.H{}})
		return
	}
	investment := helpers.GetInvestments(u)
	if investment == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "usuarios no encontrados", "data": gin.H{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "",
		"data":    investment,
	})
	return
}

func GetInvestment(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	investment := repository.GetInvestmentByID(id)
	if investment == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Usuario no existe", "data": gin.H{}})
		return
	}
	response := &httpmodels.GetInvestmentResponse{
		Investment: repository.InvestmentConvertToResponse(investment),
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "",
		"data":    response,
	})
	return

}

func UpdateInvestment(c *gin.Context) {
	u := &httpmodels.UpdateInvestmentRequest{}
	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Datos Incompletos", "data": gin.H{}})
		return
	}
	investment, errUpdate := helpers.UpdateInvestment(u)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": errUpdate.Error(), "data": gin.H{}})
		return
	}

	response := &httpmodels.UpdateInvestmentResponse{
		Investment: repository.InvestmentConvertToResponse(investment),
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Datos actualizados exitosamente",
		"data":    response,
	})
	return

}

func DeleteInvestment(c *gin.Context) {
	params := &httpmodels.DeleteRequest{}
	if err := c.Bind(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Datos Incompletos", "data": gin.H{}})
		return
	}
	err := helpers.DeleteInvestment(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Usuario no borrado", "data": gin.H{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "ok",
		"data":    gin.H{},
	})
	return
}

func CreateInvestment(c *gin.Context) {
	user := c.Value("user").(*models.Users)
	u := &httpmodels.CreateInvestmentRequest{}

	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Datos Incompletos", "data": gin.H{}})
		return
	}
	investment, errCreate := helpers.CreateInvestment(u, user)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": errCreate.Error(), "data": gin.H{}})
		return
	}

	response := &httpmodels.CreateInvestmentResponse{
		Investment: repository.InvestmentConvertToResponse(investment),
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "",
		"data":    response,
	})
	return

}
