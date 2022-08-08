package repositories

import (
	gormModels "chuck-jokes/pkg/database/models/gorm"
	"log"
	"time"

	"gorm.io/gorm"
)

// Joke base joke repository
type Joke struct {
	db *gorm.DB
}

// NewJoke create new joke repository
func NewJoke(db *gorm.DB) *Joke {
	return &Joke{
		db: db,
	}
}

// JokeOfTheDay get joke of the day from specyfi day
func (j *Joke) JokeOfTheDay(time string) *gormModels.Joke {
	var joke = gormModels.Joke{}
	j.db.Where("created_at >= ?", time).First(&joke)

	if joke.ExternalID == "" {
		return nil
	}
	return &joke
}

// GetJokes get all jokes
func (j *Joke) GetJokes(page, perPage int) *Pagination[gormModels.Joke] {
	var totalRows int64
	var jokes []gormModels.Joke
	var pagination = NewPagination[gormModels.Joke]()

	j.db.Model([]gormModels.Joke{}).Count(&totalRows)
	pagination.UpdateSettings(page, perPage)
	j.db.Scopes(paginate(pagination)).Find(&jokes)

	return pagination.PopulateData(totalRows, jokes)
}

// JokeExistInLastMonth check if same joke exist in database ,and it is newer then month
func (j *Joke) JokeExistInLastMonth(joke *gormModels.Joke) bool {
	monthEarlier := time.Now().AddDate(0, -1, 0)

	r := j.db.
		Where("id = ? AND created_at > ?", joke.ID, monthEarlier.String()).Limit(1).Find(&joke)
	if r.Error != nil {
		log.Println(r.Error)
	}

	return r.RowsAffected > 0
}

func (j *Joke) GetFavourites(page, perPage int, userID uint) *Pagination[gormModels.Joke] {
	var totalRows int64
	var jokes []gormModels.Joke
	var pagination = NewPagination[gormModels.Joke]()
	var user = gormModels.User{}

	j.db.First(&user, userID)

	totalRows = j.db.Model(&user).Association("Favourites").Count()
	pagination.UpdateSettings(page, perPage)
	dbError := j.db.Model(&user).Scopes(paginate(pagination)).Association("Favourites").Find(&jokes)
	if dbError != nil {
		log.Println(dbError)
		return nil
	}

	return pagination.PopulateData(totalRows, jokes)
}
