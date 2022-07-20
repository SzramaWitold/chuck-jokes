package database

import (
	"chuck-jokes/di"
	"database/sql"
	"fmt"
	"os"

	"gorm.io/gorm"
)

// GORMMigrator use gorm for migrate models
type GORMMigrator struct {
	db     *gorm.DB
	models []interface{}
}

// MigrateModels migrate GORM models
func (gm *GORMMigrator) MigrateModels() {
	err := gm.db.AutoMigrate(gm.models...)
	if err != nil {
		panic(err)
	}
	fmt.Println("Migration complete")
}

// NewGORMMigrator Initialize new migrator
func NewGORMMigrator(db *gorm.DB) *GORMMigrator {
	return &GORMMigrator{
		db:     db,
		models: GetAllModels(),
	}
}

var container *di.Container

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
func OpenConnection() *gorm.DB {
	if container == nil {
		newContainer := di.NewContainer()
		container = &newContainer
		fmt.Println("Database connected")
	}

	return container.Gorm
}
