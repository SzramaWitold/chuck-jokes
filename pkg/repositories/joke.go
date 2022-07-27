package repositories

import (
	gormModels "chuck-jokes/pkg/database/models/gorm"
	"log"
	"time"

	"gorm.io/gorm"
)

// JokeRepository base joke repository
type JokeRepository struct {
	db *gorm.DB
}

// NewJokeRepository create new joke repository
func NewJokeRepository(db *gorm.DB) *JokeRepository {
	return &JokeRepository{
		db: db,
	}
}

// JokeOfTheDay get joke of the day from specyfi day
func (j *JokeRepository) JokeOfTheDay(time string) *gormModels.Joke {
	var joke = gormModels.Joke{}
	j.db.Where("created_at >= ?", time).First(&joke)

	if joke.ExternalID == "" {
		return nil
	}
	return &joke
}

// GetJokes get all jokes
func (j *JokeRepository) GetJokes(page, perPage int) *Pagination {
	var totalRows int64
	var jokes = []gormModels.Joke{}
	var pagination = Pagination{}
	
	j.db.Model([]gormModels.Joke{}).Count(&totalRows)
	pagination.UpdateSettings(page, perPage)
	j.db.Scopes(paginate(&pagination)).Find(&jokes)

	return pagination.PopulateData(totalRows, jokes)
}

// JokeExistInLastMonth check if same joke exist in database and it is newer then month
func (j *JokeRepository) JokeExistInLastMonth(joke *gormModels.Joke) bool {
	monthElier := time.Now().AddDate(0, -1, 0)

	r := j.db.
		Where("id = ? AND created_at > ?", joke.ID, monthElier.String()).Limit(1).Find(&joke)
	if r.Error != nil {
		log.Println(r.Error)
	}

	return r.RowsAffected > 0
}
