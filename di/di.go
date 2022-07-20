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
	Gorm *gorm.DB
}

// NewContainer Create new conatiner with dependencies
func NewContainer() Container {
	database, err := gorm.Open(mysql.New(mysql.Config{
		DSN: getDSN(),
	}))

	if err != nil {
		panic(err)
	}

	return Container{
		Gorm: database,
	}
}

// getDSN base on .env file
func getDSN() string {
	dns := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	return dns
}
