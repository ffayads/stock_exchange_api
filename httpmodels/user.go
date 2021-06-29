package httpmodels

import (
	"time"
)

type LoginRequest struct {
	Credential string `json:"credential"`
	Password   string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type CreateUserRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type CreateUserResponse struct {
	User *User `json:"user"`
}

type UpdateUserRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Status      int    `json:"status"`
}

type UpdateUserResponse struct {
	User *User `json:"user"`
}

type GetUserResponse struct {
	User *User `json:"user"`
}

type User struct {
	ID          uint      `json:"id"`
	Status      int       `json:"status"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type PaginatorUsers struct {
	TotalRecord int64       `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Records     interface{} `json:"users"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
	Message     string      `json:"message"`
}
