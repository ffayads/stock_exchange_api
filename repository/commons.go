package repository

import (
	"github.com/api/stock_exchange_api/config/db"
	"github.com/api/stock_exchange_api/models"
)

func InitDB() {
	db.InitDB()
}

func RunMigrations() {
	db.DB.AutoMigrate(models.Users{})
	db.DB.AutoMigrate(models.Currencies{})
	db.DB.AutoMigrate(models.Companies{})
	if !db.DB.Migrator().HasConstraint(models.Companies{}, "Currencies") {
		db.DB.Migrator().CreateConstraint(models.Companies{}, "Currencies")
	}
	db.DB.AutoMigrate(models.CompaniesShare{})
	if !db.DB.Migrator().HasConstraint(models.CompaniesShare{}, "Companies") {
		db.DB.Migrator().CreateConstraint(models.CompaniesShare{}, "Companies")
	}
	db.DB.AutoMigrate(models.Instruments{})
	db.DB.AutoMigrate(models.Investments{})
	if !db.DB.Migrator().HasConstraint(models.Investments{}, "Users") {
		db.DB.Migrator().CreateConstraint(models.Investments{}, "Users")
	}
	if !db.DB.Migrator().HasConstraint(models.Investments{}, "Currencies") {
		db.DB.Migrator().CreateConstraint(models.Investments{}, "Currencies")
	}
	if !db.DB.Migrator().HasConstraint(models.Investments{}, "Companies") {
		db.DB.Migrator().CreateConstraint(models.Investments{}, "Companies")
	}
	if !db.DB.Migrator().HasConstraint(models.Investments{}, "Instruments") {
		db.DB.Migrator().CreateConstraint(models.Investments{}, "Instruments")
	}
	db.DB.AutoMigrate(models.Records{})
	if !db.DB.Migrator().HasConstraint(models.Records{}, "Investments") {
		db.DB.Migrator().CreateConstraint(models.Records{}, "Investments")
	}
}
