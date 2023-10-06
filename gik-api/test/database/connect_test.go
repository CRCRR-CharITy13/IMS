package database

import (
	"GIK_Web/database"
	"GIK_Web/env"
	"testing"

	"gorm.io/gorm"
)

var Database *gorm.DB

func TestConnectDatabase(t *testing.T) {
	env.IsLocalDB = true
	// create the test database in the same folder
	env.SqliteURI = "./test-gik-ims-localdb.sqlite"
	env.SkipMigrations = false
	database.ConnectDatabase()
	// TODO: add testing function here

	// Note: by the end, make sure to clear table

}
