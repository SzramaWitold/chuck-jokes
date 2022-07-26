package repositories

import (
	"chuck-jokes/di"
	"chuck-jokes/pkg/database"
	"fmt"
	"time"
)

// JokeOfTheDay get joke of the day from specyfi day
func JokeOfTheDay(time string) *database.Joke {
	db := di.GORM()
	var joke = database.Joke{}
	db.Where("created_at >= ?", time).First(&joke)

	if joke.ID == "" {
		return nil
	}
	return &joke
}

// GetJokes get all jokes
func GetJokes(page, perPage int) *Pagination {
	db := di.GORM()
	var totalRows int64
	var jokes = []database.Joke{}
	var pagination = Pagination{}
	
	db.Model([]database.Joke{}).Count(&totalRows)
	pagination.UpdateSettings(page, perPage)
	db.Scopes(paginate(&pagination)).Find(&jokes)

	return pagination.PopulateData(totalRows, jokes)
}

// JokeExistInLastMonth check if same joke exist in database and it is newer then month
func JokeExistInLastMonth(joke *database.Joke) bool {
	monthElier := time.Now().AddDate(0, -1, 0)

	r := di.GORM().
		Where("id = ? AND created_at > ?", joke.JokeResponse.ID, monthElier.String()).Limit(1).Find(&joke)
	if r.Error != nil {
		fmt.Println(r.Error)
	}

	return r.RowsAffected > 0
}
