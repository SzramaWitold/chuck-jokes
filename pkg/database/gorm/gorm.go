package gorm

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
		db: db,
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

// CreateDatabase with provided credentials
func CreateDatabase(dbType, dbName, user, password, host, port string) {
	dataSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/",
		user,
		password,
		host,
		port,
	)
	db, err := sql.Open(dbType, dataSource)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		dbCloseError := db.Close()
		if dbCloseError != nil {
			panic(dbCloseError)
		}
	}(db)

	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		existError := fmt.Sprintf("Error 1007: Can't create database '%v'; database exists", dbName)
		if err.Error() == existError {
			log.Println("Database:", dbName, "already exist.")
			os.Exit(1)
		} else {
			panic(err)
		}
	}

	log.Println("Database:", dbName, "created.")
}
