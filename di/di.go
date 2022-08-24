package di

import (
	"chuck-jokes/pkg/token"
	"chuck-jokes/pkg/validator"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"

	// required for myslq connection
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// container for dependencies
type container struct {
	gorm      *gorm.DB
	scheduler *gocron.Scheduler
	jwt       token.IHandler
	validator validator.IValidator
}

var cont = &container{}

func VALIDATOR(db *gorm.DB) validator.IValidator {
	if cont.validator == nil {
		cont.validator = validator.NewValidator(db)
	}

	return cont.validator
}

func JWT() *token.IHandler {
	if cont.jwt == nil {
		cont.jwt = token.NewHandler(
			os.Getenv("SECRET"),
			mustGetIntegerEnvironmentValue(os.Getenv("TTL"), 5),
			mustGetIntegerEnvironmentValue(os.Getenv("REFRESH_TTL"), 15))
	}

	return &cont.jwt
}

// GORM get gorm db connection
func GORM() *gorm.DB {
	if cont.gorm == nil {

		cont.gorm = openConnection(
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"))
	}

	return cont.gorm
}

// Scheduler go crone scheduler connection
func Scheduler() *gocron.Scheduler {
	if cont.scheduler == nil {
		cont.scheduler = gocron.NewScheduler(time.UTC)
	}

	return cont.scheduler
}

//OpenConnection for database inside DB var
func openConnection(user, password, host, port, name string) *gorm.DB {
	if cont.gorm == nil {
		database, err := gorm.Open(mysql.New(mysql.Config{
			DSN: getDSN(user, password, host, port, name),
		}))

		if err != nil {
			panic(err)
		}
		cont.gorm = database
		log.Println("Database connected")
	}

	return cont.gorm
}

// getDSN base on .env file
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
