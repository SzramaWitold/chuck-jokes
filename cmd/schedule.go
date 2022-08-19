package cmd

import (
	"chuck-jokes/di"
	"chuck-jokes/models"
	gormModels "chuck-jokes/pkg/database/gorm/models"
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/requests"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(schedulerCmd)
}

var schedulerCmd = &cobra.Command{
	Use:   "schedule:run",
	Short: "schedule all job inside crone",
	Long:  `Schedule everything inside crone package`,
	Run: func(_ *cobra.Command, _ []string) {

		gorm := di.GORM(
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		scheduler := di.Scheduler()
		callSchedules(scheduler, gorm)
		scheduler.StartBlocking()
	},
}

func callSchedules(scheduler *gocron.Scheduler, db *gorm.DB) {
	external := requests.NewExternalRequest(os.Getenv("EXTERNAL_API") + "jokes/random")
	_, sRanJokeErr := scheduler.Every(1).Minute().Do(scheduleRandomJoke(db, external.CallRandom))
	if sRanJokeErr != nil {
		log.Panic(sRanJokeErr)
	}
}

func scheduleRandomJoke(db *gorm.DB, external func() *models.Joke) func() {
	return func() {
		JokeRepository := repositories.NewJoke(db)
		joke := external()
		if JokeRepository.JokeExistInLastMonth(joke) {
			scheduleRandomJoke(db, external)
			return
		}

		gormModels.Create(db, joke)

		log.Println("Random joke scheduler finished")
	}
}
