package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/api/stock_exchange_api/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func CreateTokenUser(user *models.Users) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_email"] = user.Email
	claims["user_id"] = fmt.Sprint(user.ID)
	claims["token_type"] = "user"
	claims["exp"] = time.Now().Add(time.Hour * 180).Unix()
	claims["iat"] = time.Now().Unix()
	fmt.Println(claims)
	// Generate encoded token and send it as web.
	var JWTKEY = os.Getenv("JWTKEY")
	t, err := token.SignedString([]byte(JWTKEY))
	if err != nil {
		log.Println("", err)
		return "", err
	}
	return t, nil
}

func GetToken(s string) (*jwt.Token, error) {
	key := os.Getenv("JWTKEY")
	// parsing token
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Token invalido"))
	}
	return token, nil
}

func GetTokenByField(token, field string) (string, error) {
	tokenU, err := GetToken(token)
	if err != nil {
		return "", err
	}
	var response string
	if claimsU, ok := tokenU.Claims.(jwt.MapClaims); ok && tokenU.Valid {
		if claimsU[field] == nil {
			return "", nil
		}
		response = claimsU[field].(string)
	} else {
		return "", errors.New("Toke invalido")
	}
	return response, nil
}
