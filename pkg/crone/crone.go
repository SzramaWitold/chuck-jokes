package crone

import (
	models "chuck-jokes/pkg/database/models/gorm"
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/requests"
	"gorm.io/gorm"
	"log"

	"github.com/go-co-op/gocron"
)

// CronScheduler base struct
type CronScheduler struct {
	scheduler *gocron.Scheduler
	db        *gorm.DB
}

// NewCronScheduler create new crone scheduler
func NewCronScheduler(scheduler *gocron.Scheduler, db *gorm.DB) *CronScheduler {
	return &CronScheduler{
		scheduler: scheduler,
		db:        db,
	}
}

func (c *CronScheduler) scheduleRandomJoke() {
	JokeRepository := repositories.NewJoke(c.db)
	joke := requests.CallRandom()
	dbJoke := models.ChangeToGormJoke(&joke)

	if JokeRepository.JokeExistInLastMonth(&dbJoke) {
		c.scheduleRandomJoke()
	}

	dbJoke.Create(c.db)

	log.Println("Scheduler run")
}

// Schedule all
func (c *CronScheduler) Schedule(async bool) {
	_, sRanJokeErr := c.scheduler.Every(1).Minute().Do(c.scheduleRandomJoke)
	if sRanJokeErr != nil {
		log.Panic(sRanJokeErr)
	}

	if async {
		c.scheduler.StartAsync()
	} else {
		c.scheduler.StartBlocking()
	}

}
