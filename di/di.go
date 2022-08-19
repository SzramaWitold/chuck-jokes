package di

import (
	"chuck-jokes/pkg/token"
	"chuck-jokes/pkg/validator"
	"fmt"
	"log"
	"strconv"
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
	jwt       token.IHandler
	validator *validator.Validator
}

var container = &Container{}

func VALIDATOR() *validator.Validator {
	if container.validator == nil {
		container.validator = validator.NewValidator()
	}

	return container.validator
}

func JWT(secret, ttl, refreshTtl string) *token.IHandler {
	if container.jwt == nil {
		ttlDuration, refreshTTlDuration := setTTLAndRefresh(ttl, refreshTtl)
		container.jwt = token.NewHandler(secret, ttlDuration, refreshTTlDuration)
	}

	return &container.jwt
}

// GORM get gorm db connection
func GORM(user, password, host, port, name string) *gorm.DB {
	if container.gorm == nil {
		container.gorm = openConnection(user, password, host, port, name)
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
func openConnection(user, password, host, port, name string) *gorm.DB {
	if container.gorm == nil {
		database, err := gorm.Open(mysql.New(mysql.Config{
			DSN: getDSN(user, password, host, port, name),
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

func setTTLAndRefresh(ttl, refreshTtl string) (int, int) {
	ttlDuration := 5
	refreshTTLDuration := 10
	if ttl != "" {
		newTTL, ttlErr := strconv.Atoi(ttl)
		if ttlErr != nil {
			log.Println(ttlErr)
		} else {
			ttlDuration = newTTL
		}
	}

	if ttl != "" {
		newRefreshTTL, ttlErr := strconv.Atoi(refreshTtl)
		if ttlErr != nil {
			log.Println(ttlErr)
		} else {
			refreshTTLDuration = newRefreshTTL
		}
	}

	return ttlDuration, refreshTTLDuration
}
