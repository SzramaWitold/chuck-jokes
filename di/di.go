package di

import (
	"chuck-jokes/pkg/token"
	"chuck-jokes/pkg/validator"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"

	// required for myslq connection
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Container for dependencies
type Container struct {
	gorm      *gorm.DB
	scheduler *gocron.Scheduler
	jwt       *token.Handler
	validator *validator.Validator
}

var container = &Container{}

func VALIDATOR() *validator.Validator {
	if container.validator == nil {
		container.validator = validator.NewValidator()
	}

	return container.validator
}

func JWT() *token.Handler {
	if container.jwt == nil {
		container.jwt = token.NewHandler()
	}

	return container.jwt
}

// GORM get gorm db connection
func GORM() *gorm.DB {
	if container.gorm == nil {
		container.gorm = openConnection()
	}

	return container.gorm
}

// Scheduler go crone scheduler connection
func Scheduler() *gocron.Scheduler {
	if container.scheduler == nil {
		container.scheduler = gocron.NewScheduler(time.UTC)
	}

	return container.scheduler
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
		log.Println("Database connected")
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
