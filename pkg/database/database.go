package database

import (
	"database/sql"
	"fmt"
	"os"

	// required for myslq connection
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB Initialize FORM database
var DB *gorm.DB

// GetDNS base on .env file
func GetDNS() string {
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

// CreateDatabase base on .env file
func CreateDatabase() {
	dataSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	db, err := sql.Open(os.Getenv("DB_TYPE"), dataSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE " + os.Getenv("DB_NAME"))
	if err != nil {
		existError := fmt.Sprintf("Error 1007: Can't create database '%v'; database exists", os.Getenv("DB_NAME"))
		if err.Error() == existError {
			fmt.Println("Database:", os.Getenv("DB_NAME"), "already exist.")
			os.Exit(1)
		} else {
			panic(err)
		}
	}

	fmt.Println("Database:", os.Getenv("DB_NAME"), "created.")
}

//OpenConnection for database inside DB var
func OpenConnection() {
	database, err := gorm.Open(mysql.New(mysql.Config{
		DSN: GetDNS(),
	}))

	if err != nil {
		panic(err)
	}
	DB = database
	fmt.Println("Database connected")
}

// MigrateModels migrate GORM models
func MigrateModels(models []interface{}) {
	OpenConnection()

	DB.AutoMigrate(models...)
}
