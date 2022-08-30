package di

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// OpenConnection for database inside DB var.
func openConnection(user, password, host, port, name string) *gorm.DB {
	database, err := gorm.Open(mysql.New(mysql.Config{
		DSN: getDSN(user, password, host, port, name),
	}))
	if err != nil {
		panic(err)
	}

	cont.gorm = database

	log.Println("Database connected")

	return cont.gorm
}

// getDSN base on .env file.
func getDSN(user, password, host, port, name string) string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		name,
	)
}

func mustGetIntegerEnvironmentValue(val string, def int) int {
	if val != "" {
		intVal, ttlErr := strconv.Atoi(val)
		if ttlErr != nil {
			log.Println(ttlErr)

			return def
		}

		return intVal
	} else {
		return def
	}
}
