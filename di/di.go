package di

import (
	"os"
	"time"

	"chuck-jokes/pkg/token"

	"github.com/go-co-op/gocron"
	"github.com/go-playground/validator/v10"
	// required for myslq connection.
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// container for dependencies.
type container struct {
	gorm      *gorm.DB
	scheduler *gocron.Scheduler
	jwt       token.TokenHandler
	validator *validator.Validate
}

var cont = &container{}

func Validator() *validator.Validate {
	if cont.validator == nil {
		cont.validator = validator.New()
	}

	return cont.validator
}

func JWT() *token.TokenHandler {
	if cont.jwt == nil {
		cont.jwt = token.NewHandler(
			os.Getenv("SECRET"),
			mustGetIntegerEnvironmentValue(os.Getenv("TTL"), 5),
			mustGetIntegerEnvironmentValue(os.Getenv("REFRESH_TTL"), 15))
	}

	return &cont.jwt
}

// GORM get gorm db connection.
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

// Scheduler go crone scheduler connection.
func Scheduler() *gocron.Scheduler {
	if cont.scheduler == nil {
		cont.scheduler = gocron.NewScheduler(time.UTC)
	}

	return cont.scheduler
}
