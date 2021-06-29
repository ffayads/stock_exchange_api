package repository

import (
	"github.com/api/stock_exchange_api/config/db"
	"github.com/api/stock_exchange_api/httpmodels"
	"github.com/api/stock_exchange_api/models"
	"github.com/api/stock_exchange_api/utils"
)

func UsersConvertToListResponse(users []models.Users) []*httpmodels.User {
	usersResp := []*httpmodels.User{}
	for _, user := range users {
		usersResp = append(usersResp, UserConvertToResponse(&user))
	}
	return usersResp
}

func UserConvertToResponse(user *models.Users) *httpmodels.User {

	userRes := &httpmodels.User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Status:      user.Status,
	}
	return userRes
}

func GetUsers(params *httpmodels.Filter) []models.Users {
	users := []models.Users{}
	query := db.DB
	query = utils.Filter(query, params)
	query.Find(&users)
	return users
}

func GetUserByFilters(params *httpmodels.Filter) *models.Users {
	users := &models.Users{}
	query := db.DB
	query = utils.Filter(query, params)
	if err := query.First(users).Error; err != nil {
		return nil
	}
	return users
}

func GetUserByID(id interface{}) *models.Users {

	user := models.Users{}
	query := db.DB.Where("id = ?", id)
	if query.First(&user).Error != nil {
		return nil
	}
	return &user
}

func CreateUser(params *httpmodels.CreateUserRequest) (*models.Users, error) {
	u := &models.Users{
		Name:        params.Name,
		Email:       params.Email,
		PhoneNumber: params.PhoneNumber,
		Password:    params.Password,
		Status:      models.USER_ACTIVE,
	}
	err := u.Create()
	return u, err
}

func UpdateUser(params *httpmodels.UpdateUserRequest) (*models.Users, error) {
	u := &models.Users{}
	if db.DB.Where("id = ?", params.ID).First(u).Error != nil {
		return u, nil
	}
	u.Name = params.Name
	u.PhoneNumber = params.PhoneNumber
	u.Password = params.Password
	u.Status = params.Status
	err := u.Save()
	return u, err
}

func DeleteUser(params *httpmodels.DeleteRequest) error {
	c := models.Users{}
	if db.DB.Where("id = ?", params.ID).First(&c).Error != nil {
		return nil
	}
	err := c.Delete()
	return err
}
