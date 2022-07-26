package di

import (
	"fmt"
	"os"

	// required for myslq connection
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Container for dependencies
type Container struct {
	gorm *gorm.DB
}

var container = &Container{}

// GORM get gorm db connection
func GORM() *gorm.DB {
	if container.gorm == nil {
		container.gorm = openConnection()
	}

	return container.gorm
}

//OpenConnection for database inside DB var
func openConnection() *gorm.DB {
	if container.gorm == nil {
		database, err := gorm.Open(mysql.New(mysql.Config{
			DSN: getDSN(),
		}))

		if err != nil {
			panic(err)
		}
		container.gorm = database
		fmt.Println("Database connected")
	}

	return container.gorm
}

// getDSN base on .env file
func getDSN() string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
