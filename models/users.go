package models

import (
	"time"

	"github.com/api/stock_exchange_api/config/db"
	"gorm.io/gorm"
)

var USER_ACTIVE = 0
var USER_EXPIRED = 1
var USER_BLOCKED = 2
var USER_RESET = 3

type Users struct {
	BaseModel
	Name            string    `json:"name" gorm:"column:name"`
	Email           string    `json:"email" gorm:"column:email"`
	PhoneNumber     string    `json:"phone_number" gorm:"column:phone_number"`
	Password        string    `json:"password" gorm:"column:password"`
	PasswordUpdated time.Time `json:"password_updated" gorm:"column:password_updated"`
	Status          int       `json:"status" gorm:"column:status"`
}

type UsersArray []Users

func (u *Users) Create() error {
	return db.DB.Create(u).Error
}

func (u *Users) Save() error {
	return db.DB.Save(u).Error
}

func (u *Users) Delete() error {
	return db.DB.Delete(u).Error
}

func (u *Users) BeforeSave(tx *gorm.DB) {
	if tx.Statement.Changed("Password") {
		if u.Password != "" {
			tx.Statement.SetColumn("PasswordUpdated", time.Now())
		}
	}
}
