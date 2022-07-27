package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

// Manager use for create and migrate
type Manager struct {
	db     *gorm.DB
	models []interface{}
}

// NewManager Initialize new migrator
func NewManager(db *gorm.DB) *Manager {
	return &Manager{
		db:     db,
	}
}

// MigrateModels migrate GORM models
func (m *Manager) MigrateModels(args ...interface{}) {
	err := m.db.AutoMigrate(args...)
	if err != nil {
		panic(err)
	}
	log.Println("Migration complete")
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
			log.Println("Database:", os.Getenv("DB_NAME"), "already exist.")
			os.Exit(1)
		} else {
			panic(err)
		}
	}

	log.Println("Database:", os.Getenv("DB_NAME"), "created.")
}