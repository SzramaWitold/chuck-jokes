package cmd

import (
	"log"
	"net/http"
	"os"

	"chuck-jokes/di"
	"chuck-jokes/models"
	gormRepository "chuck-jokes/pkg/repositories/gorm"
	"chuck-jokes/pkg/requests"

	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func init() {
	rootCmd.AddCommand(schedulerCmd)
}

var schedulerCmd = &cobra.Command{
	Use:   "schedule:run",
	Short: "schedule all job inside crone",
	Long:  `Schedule everything inside crone package`,
	Run: func(_ *cobra.Command, _ []string) {
		gorm := di.GORM()
		scheduler := di.Scheduler()
		callSchedules(scheduler, gorm)
		scheduler.StartBlocking()
	},
}

func callSchedules(scheduler *gocron.Scheduler, db *gorm.DB) {
	external := requests.NewExternalRequest(os.Getenv("EXTERNAL_API")+"jokes/random", http.DefaultClient)
	_, sRanJokeErr := scheduler.Every(1).Minute().Do(scheduleRandomJoke(db, external.CallRandom))

	if sRanJokeErr != nil {
		log.Panic(sRanJokeErr)
	}
}

func scheduleRandomJoke(db *gorm.DB, external func() *models.Joke) func() {
	return func() {
		JokeRepository := gormRepository.NewJoke(db)
		joke := external()

		if JokeRepository.JokeExistInLastMonth(joke) {
			scheduleRandomJoke(db, external)

			return
		}

		if _, createErr := JokeRepository.Create(joke); createErr != nil {
			log.Println(createErr)
		}

		log.Println("Random joke scheduler finished")
	}
}
