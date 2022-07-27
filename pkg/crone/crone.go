package crone

import (
	"chuck-jokes/di"
	"chuck-jokes/pkg/database/models/gorm"
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/requests"
	"log"

	"github.com/go-co-op/gocron"
)

// CronScheduler base struct
type CronScheduler struct {
	scheduler *gocron.Scheduler
}

// NewCronScheduler create new crone scheduler
func NewCronScheduler(scheduler *gocron.Scheduler) *CronScheduler {
	return &CronScheduler{scheduler: scheduler}
}

func scheduleRandomJoke() {
	db := di.GORM()
	JokeRepository := repositories.NewJokeRepository(db)
	joke := requests.CallRandom()
	dbJoke := gorm.ChangeToGormJoke(&joke)

	if JokeRepository.JokeExistInLastMonth(&dbJoke) {
		scheduleRandomJoke()
	}

	dbJoke.Create(db)

	log.Println("Scheduler runed")
}

// Schedule all
func (c *CronScheduler) Schedule(async bool) {
	c.scheduler.Every(1).Minute().Do(scheduleRandomJoke)

	if async {
		c.scheduler.StartAsync()
	} else {
		c.scheduler.StartBlocking()
	}

}
