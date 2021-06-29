package controllers

import (
	"log"
	"net/http"

	"github.com/api/stock_exchange_api/helpers"
	"github.com/api/stock_exchange_api/httpmodels"
	"github.com/api/stock_exchange_api/repository"
	"github.com/api/stock_exchange_api/utils"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	u := &httpmodels.Filter{}
	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Datos Incompletos", "data": gin.H{}})
		return
	}
	user := helpers.GetUsers(u)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "usuarios no encontrados", "data": gin.H{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "",
		"data":    user,
	})
	return
}

func GetUser(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	user := repository.GetUserByID(id)

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Usuario no existe", "data": gin.H{}})
		return
	}

	response := &httpmodels.GetUserResponse{
		User: repository.UserConvertToResponse(user),
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "",
		"data":    response,
	})
	return

}

func UpdateUser(c *gin.Context) {
	u := &httpmodels.UpdateUserRequest{}
	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Datos Incompletos", "data": gin.H{}})
		return
	}
	user, errUpdate := helpers.UpdateUser(u)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": errUpdate.Error(), "data": gin.H{}})
		return
	}

	response := &httpmodels.UpdateUserResponse{
		User: repository.UserConvertToResponse(user),
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Datos actualizados exitosamente",
		"data":    response,
	})
	return

}

func DeleteUser(c *gin.Context) {
	params := &httpmodels.DeleteRequest{}
	if err := c.Bind(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Datos Incompletos", "data": gin.H{}})
		return
	}
	err := helpers.DeleteUser(params)
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

func UserLogin(c *gin.Context) {
	u := &httpmodels.LoginRequest{}

	token := ""
	message := ""

	response := httpmodels.LoginResponse{
		Token: token,
	}
	if err := c.Bind(u); err != nil {
		log.Println("Datos incompletos")
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Datos Incompletos", "data": response})
		return
	}

	user, errValidate := helpers.ValidateCredentialsUser(u)
	if user == nil {
		log.Println("Credenciales erroneas")
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": errValidate.Error(), "data": response})
		return
	}

	if errValidate == nil {
		var errToken error
		token, errToken = utils.CreateTokenUser(user)
		if errToken != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": errToken.Error(), "data": response})
			return
		}
		response.Token = token
		message = "Login exitoso"
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": message, "data": response})
		return
	} else {
		message = errValidate.Error()
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": message, "data": response})
		return
	}

}

func CreateUser(c *gin.Context) {
	u := &httpmodels.CreateUserRequest{}

	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Datos Incompletos", "data": gin.H{}})
		return
	}
	user, errCreate := helpers.CreateUser(u)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": errCreate.Error(), "data": gin.H{}})
		return
	}

	response := &httpmodels.CreateUserResponse{
		User: repository.UserConvertToResponse(user),
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "",
		"data":    response,
	})
	return

}
