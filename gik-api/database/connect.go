package database

import (
	"GIK_Web/env"
	"GIK_Web/types"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDatabase() {
	dsn := env.MysqlURi
	fmt.Println("DSN: " + dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database: " + err.Error())
	}

	Database = db

	migrations()
}

func migrations() {
	// Database.AutoMigrate(&models.Whatever{})

	if env.SkipMigrations {
		return
	}

	Database.AutoMigrate(&types.Item{})
	Database.AutoMigrate(&types.Tag{})
	Database.AutoMigrate(&types.User{})
	Database.AutoMigrate(&types.Client{})
	Database.AutoMigrate(&types.Transaction{})
	Database.AutoMigrate(&types.TransactionItem{})
	Database.AutoMigrate(&types.Session{})
	Database.AutoMigrate(&types.AdvancedLog{})
	Database.AutoMigrate(&types.SimpleLog{})
	Database.AutoMigrate(&types.SignupCode{})
	Database.AutoMigrate(&types.Location{})
}
