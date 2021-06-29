package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/api/stock_exchange_api/controllers"

	"github.com/api/stock_exchange_api/repository"
	"github.com/api/stock_exchange_api/utils"
)

func main() {

	repository.InitDB()
	repository.RunMigrations()

	r := gin.Default()
	r.Use(CORSMiddleware())
	v1 := r.Group("/v1")
	r.GET("/", MiddlewareRequest(false), controllers.Index)
	v1.POST("/user/login", MiddlewareRequest(false), controllers.UserLogin)
	v1.POST("/user/create", MiddlewareRequest(false), controllers.CreateUser)
	v1.PUT("/user/update", MiddlewareRequest(true), controllers.UpdateUser)
	v1.DELETE("/user/delete", MiddlewareRequest(true), controllers.DeleteUser)
	v1.POST("/investments/get", MiddlewareRequest(true), controllers.GetInvestments)
	v1.POST("/investments/create", MiddlewareRequest(true), controllers.CreateInvestment)
	v1.PUT("/investments/update", MiddlewareRequest(true), controllers.UpdateInvestment)
	v1.DELETE("/investments/delete", MiddlewareRequest(true), controllers.DeleteInvestment)
	v1.POST("/records/get", MiddlewareRequest(true), controllers.GetRecordDetail)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	r.Run(":" + port)
}

//Validacion de token
func MiddlewareRequest(verifySession bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Request.Header)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Session-Token, session_token")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		if verifySession {
			log.Println(c.Request.Header)
			sessionToken := c.Request.Header.Get("Session-Token")
			log.Println(sessionToken)
			if len(sessionToken) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Session Token no recibido", "data": gin.H{}})
				c.Abort()
				return

			}
			tokenU, err := utils.GetToken(sessionToken)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error(), "data": gin.H{}})
				c.Abort()
				return
			}
			if claimsU, ok := tokenU.Claims.(jwt.MapClaims); ok && tokenU.Valid {

				if claimsU["token_type"].(string) == "user" {
					user := repository.GetUserByID(string(claimsU["user_id"].(string)))
					if user == nil {
						c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Usuario no encontrado", "data": gin.H{}})
						c.Abort()
						return
					}

					c.Set("user", user)

				} else {
					c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Token incorrecto", "data": gin.H{}})
					c.Abort()
					return
				}

			} else {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Token incorrecto", "data": gin.H{}})
				c.Abort()
				return
			}
		}

		c.Next()

	}
}

//Headers de seguridad
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Session-Token, session_token")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
	}
}
