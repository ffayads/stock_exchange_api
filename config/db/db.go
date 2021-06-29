package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {

	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", os.Getenv("DBUSERNAME"), os.Getenv("DBUSERPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))
	fmt.Println(dbConfig)
	DB, err = gorm.Open(mysql.Open(dbConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	fmt.Println(DB)

	if err != nil {
		panic(err)
	}
}
