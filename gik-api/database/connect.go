package database

import (
	"GIK_Web/env"
	"GIK_Web/types"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDatabase() {
	if !env.IsLocalDB {
		dsn := env.MysqlURi
		fmt.Println("DSN: " + dsn)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Unable to connect to database: " + err.Error())
			// try creating the database, if skipMigrations == false
			if !env.SkipMigrations {
				migrations()
				db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			} else {
				return
			}
		}
		Database = db

	} else {
		fmt.Printf("Connect to the local database: %s \n", env.SqliteURI)
		db, err := gorm.Open(sqlite.Open(env.SqliteURI), &gorm.Config{})
		if err != nil {
			fmt.Println("Error: ", err)
			if !env.SkipMigrations {
				migrations()
				db, err = gorm.Open(sqlite.Open(env.SqliteURI), &gorm.Config{})
			} else {
				return
			}
		}
		Database = db
	}
	migrations()
}

func migrations() {
	// Database.AutoMigrate(&models.Whatever{})

	if env.SkipMigrations {
		return
	}
	Database.AutoMigrate(&types.Item{})
	Database.AutoMigrate(&types.Location{})
	Database.AutoMigrate(&types.Warehouse{})
	Database.AutoMigrate(&types.User{})
	Database.AutoMigrate(&types.Donor{})
	Database.AutoMigrate(&types.Client{})
	Database.AutoMigrate(&types.Donation{})
	Database.AutoMigrate(&types.DonationItem{})
	Database.AutoMigrate(&types.Order{})
	Database.AutoMigrate(&types.OrderItem{})
	Database.AutoMigrate(&types.Session{})
	Database.AutoMigrate(&types.SignupCode{})
	Database.AutoMigrate(&types.AdvancedLog{})
	Database.AutoMigrate(&types.SimpleLog{})
}
