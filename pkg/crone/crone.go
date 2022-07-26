package crone

import (
	"chuck-jokes/pkg/database"
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/requests"
	"fmt"
	"github.com/go-co-op/gocron"
)

type CronScheduler struct {
	scheduler *gocron.Scheduler
}

func NewCronScheduler(scheduler *gocron.Scheduler) *CronScheduler {
	return &CronScheduler{scheduler: scheduler}
}

func scheduleRandomJoke() {
	joke := requests.CallRandom()
	dbJoke := database.Joke{JokeResponse: joke}

	if repositories.JokeExistInLastMonth(&dbJoke) {
		scheduleRandomJoke()
	}

	database.CreateJoke(&database.Joke{JokeResponse: joke})
	fmt.Println("Scheduler runed")
}

// Schedule all
func (s *CronScheduler) Schedule(async bool) {
	s.scheduler.Every(1).Minute().Do(scheduleRandomJoke)

	if async {
		s.scheduler.StartAsync()
	} else {
		s.scheduler.StartBlocking()
	}
}
