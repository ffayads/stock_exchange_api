package models

import (
	"time"

	//"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// DB is a connection to your database to be used
// throughout your application.
var DBManager *gorm.DB

type BaseModel struct {
	//ID uuid.UUID `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4();column:id"`
	ID        uint       `json:"id" sql:"primary_key;column:id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

/*
func (b *BaseModel) BeforeCreate(scope *gorm.DB) error {
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return err
	}
	b.ID = id
	return nil
}
*/
