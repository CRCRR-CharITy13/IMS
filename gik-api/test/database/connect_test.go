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
	env.SqliteURI = "./assets/test-gik-ims-localdb.sqlite"
	env.SkipMigrations = false
	database.ConnectDatabase()
	// TODO: error checking
}
