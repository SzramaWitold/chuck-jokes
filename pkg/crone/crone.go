package crone

import (
	"chuck-jokes/pkg/database"
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/requests"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

var scheduler *gocron.Scheduler

// GetScheduler or create new scheduler
func GetScheduler() *gocron.Scheduler {
	if scheduler == nil {
		scheduler = gocron.NewScheduler(time.UTC)
		fmt.Println("Scheduler created")
	}

	return scheduler
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
func Schedule(async bool) {
	GetScheduler().Every(1).Minute().Do(scheduleRandomJoke)

	if async {
		GetScheduler().StartAsync()
	} else {
		GetScheduler().StartBlocking()
	}

}
