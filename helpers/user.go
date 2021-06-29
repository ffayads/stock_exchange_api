package helpers

import (
	"errors"
	"fmt"
	"time"

	"github.com/api/stock_exchange_api/httpmodels"
	"github.com/api/stock_exchange_api/models"
	"github.com/api/stock_exchange_api/repository"
	"github.com/api/stock_exchange_api/utils"
)

func ValidateCredentialsUser(u *httpmodels.LoginRequest) (*models.Users, error) {
	u.Credential = utils.CleanEmail(u.Credential)
	field := "email"
	filter := utils.CreateFilter(field, u.Credential)
	user := repository.GetUserByFilters(filter)
	if user == nil {
		return nil, errors.New("Usuario no encontrado")
	}
	days := time.Now().Sub(user.PasswordUpdated).Hours() / 24
	if days >= 90 {
		user.Status = models.USER_EXPIRED
	}
	fmt.Println("Password ", user.Password)
	fmt.Println("Password 2 ", u.Password)
	err := utils.ComparePassword(user.Password, u.Password)
	fmt.Println("errrorrr ", err)
	if err != nil {
		return user, errors.New("Contraseña invalida")
	}
	return user, nil
}

func CreateUser(u *httpmodels.CreateUserRequest) (*models.Users, error) {
	u.Email = utils.CleanEmail(u.Email)
	u.PhoneNumber = utils.CleanPhoneNumber(u.PhoneNumber)
	field := "email"
	filter := utils.CreateFilter(field, u.Email)
	if repository.GetUserByFilters(filter) != nil {
		return nil, errors.New("Usuario ya existe")
	}

	errV := utils.ValidatePasswordPolices(u.Password)
	if errV != nil {
		fmt.Println(errV)
		return nil, errV
	}
	password, errG := utils.GeneratePassword(u.Password)
	if errG != nil {
		return nil, errors.New("Contraseña no valida")
	}
	u.Password = password

	user, err := repository.CreateUser(u)
	if err != nil {
		return nil, errors.New("No se pudo crear el usuario")
	}
	return user, nil
}

func UpdateUser(u *httpmodels.UpdateUserRequest) (*models.Users, error) {
	phoneNumber := utils.CleanPhoneNumber(u.PhoneNumber)
	if phoneNumber != "" {
		u.PhoneNumber = phoneNumber
	}
	errV := utils.ValidatePasswordPolices(u.Password)
	if errV != nil {
		fmt.Println(errV)
		return nil, errV
	}
	password, errG := utils.GeneratePassword(u.Password)
	if errG != nil {
		return nil, errors.New("Contraseña no valida")
	}
	u.Password = password
	user, err := repository.UpdateUser(u)
	return user, err
}

func DeleteUser(params *httpmodels.DeleteRequest) error {
	err := repository.DeleteUser(params)
	return err
}

func GetUsers(values *httpmodels.Filter) []models.Users {
	return repository.GetUsers(values)
}
