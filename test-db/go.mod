module db-sqlite-test

go 1.18

require (
	gorm.io/driver/sqlite v1.5.2
	gorm.io/gorm v1.25.3
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	test-db/data/data-utils v0.0.0-00010101000000-000000000000 // indirect
)

replace test-db/data/data-utils => ./data/
