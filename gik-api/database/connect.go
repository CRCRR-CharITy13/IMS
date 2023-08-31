package database

import (
	"GIK_Web/env"
	"GIK_Web/type_news"
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
			panic("Unable to connect to database: " + err.Error())
		}
		Database = db

	} else {
		fmt.Printf("Connect to the local database: %s", env.SqliteURI)
		db, err := gorm.Open(sqlite.Open(env.SqliteURI), &gorm.Config{})
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		//

		// Get a reference to the underlying database connection
		// sqlDB, err := db.DB()
		// if err != nil {
		// 	panic("failed to get database connection")
		// }
		// defer sqlDB.Close()

		// // Query the database for table names
		// rows, err := sqlDB.Query("SELECT name FROM sqlite_master WHERE type='table'")
		// if err != nil {
		// 	panic("failed to query database")
		// }
		// defer rows.Close()

		// fmt.Println("Tables:")
		// for rows.Next() {
		// 	var tableName string
		// 	if err := rows.Scan(&tableName); err != nil {
		// 		panic("failed to scan row")
		// 	}
		// 	fmt.Println(tableName)
		// }

		//
		Database = db
	}
	//migrations()
}

func migrations() {
	// Database.AutoMigrate(&models.Whatever{})

	if env.SkipMigrations {
		return
	}
	if !env.IsLocalDB {
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
	} else {
		Database.AutoMigrate(&type_news.Item{})
		Database.AutoMigrate(&type_news.Location{})
		Database.AutoMigrate(&type_news.Warehouse{})
		Database.AutoMigrate(&type_news.User{})
		Database.AutoMigrate(&type_news.Donor{})
		Database.AutoMigrate(&type_news.Client{})
		Database.AutoMigrate(&type_news.Donation{})
		Database.AutoMigrate(&type_news.DonationItem{})
		Database.AutoMigrate(&type_news.Order{})
		Database.AutoMigrate(&type_news.OrderItem{})
		Database.AutoMigrate(&type_news.Session{})
		Database.AutoMigrate(&type_news.AdvancedLog{})
		Database.AutoMigrate(&type_news.SimpleLog{})
		Database.AutoMigrate(&type_news.Tag{})
	}
}
