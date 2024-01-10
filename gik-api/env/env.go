package env

import (
	"os"

	"github.com/joho/godotenv"
)

var WebserverPort string

var DebugMode bool

var MysqlURi string

var SkipMigrations bool

var HTTPS bool

var JWTSigningPassword string

var IsLocalDB bool

var SqliteURI string

// add the env_path to create the testing environment
func SetEnv() {
	//
	godotenv.Load(".env")

	WebserverPort = os.Getenv("PORT")
	if WebserverPort == "" {
		panic("HOST and PORT are not set")
	}

	SkipMigrations = os.Getenv("SKIP_MIGRATIONS") == "true"

	DebugMode = os.Getenv("DEBUG_MODE") == "true"

	HTTPS = os.Getenv("HTTPS") == "true"

	MysqlURi = os.Getenv("MYSQL_URI")
	if MysqlURi == "" {
		panic("MYSQL_URI is not set")
	}

	JWTSigningPassword = os.Getenv("JWT_SIGNING_PASSWORD")
	if JWTSigningPassword == "" {
		panic("JWT_SIGNING_PASSWORD is not set")
	}

	IsLocalDB = os.Getenv("LOCAL_DB") == "true"

	SqliteURI = os.Getenv("SQLITE_URI")
}
