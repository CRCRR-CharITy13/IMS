package env

import (
	"os"

	"github.com/joho/godotenv"
)

var WebserverHost string
var WebserverPort string

var DebugMode bool

var MysqlURi string

var SkipMigrations bool

var HTTPS bool

var CookieDomain string

var JWTSigningPassword string

func SetEnv() {
	godotenv.Load(".env")

	WebserverHost = os.Getenv("HOST")
	WebserverPort = os.Getenv("PORT")
	if WebserverHost == "" || WebserverPort == "" {
		panic("HOST and PORT are not set")
	}

	SkipMigrations = os.Getenv("SKIP_MIGRATIONS") == "true"

	DebugMode = os.Getenv("DEBUG_MODE") == "true"

	HTTPS = os.Getenv("HTTPS") == "true"

	MysqlURi = os.Getenv("MYSQL_URI")
	if MysqlURi == "" {
		panic("MYSQL_URI is not set")
	}

	CookieDomain = os.Getenv("COOKIE_DOMAIN")
	if CookieDomain == "" {
		panic("COOKIE_DOMAIN is not set")
	}

	JWTSigningPassword = os.Getenv("JWT_SIGNING_PASSWORD")
	if JWTSigningPassword == "" {
		panic("JWT_SIGNING_PASSWORD is not set")
	}
}
