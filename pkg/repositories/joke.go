package repositories

import (
	"chuck-jokes/models"
	gormModels "chuck-jokes/pkg/database/gorm/models"
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
func (j *Joke) JokeOfTheDay(time string) *models.Joke {
	var joke = models.Joke{}
	j.db.Where("created_at >= ?", time).First(&joke)

	if joke.ExternalID == "" {
		return nil
	}
	return &joke
}

// GetJoke get joke of the day from specyfi day
func (j *Joke) GetJoke(jokeID uint) *models.Joke {
	var joke = models.Joke{}
	txFind := j.db.First(&joke, jokeID)

	if txFind.Error != nil {
		log.Println(txFind.Error)
		return nil
	}

	joke.Shows++

	txSave := j.db.Save(&joke)

	if txSave.Error != nil {
		log.Println(txSave.Error)
		return nil
	}

	return &joke
}

// GetJoke get joke of the day from specyfi day
func (j *Joke) GetStatistic(jokeID uint) (*models.Joke, uint) {
	var joke = models.Joke{}
	txFind := j.db.First(&joke, jokeID)

	if txFind.Error != nil {
		log.Println(txFind.Error)
		return nil, 0
	}

	favourites := j.db.Model(&joke).Association("Users").Count()

	return &joke, uint(favourites)
}

// GetJokes get all jokes
func (j *Joke) GetJokes(page, perPage int) *Pagination[models.Joke] {
	var totalRows int64
	var jokes []models.Joke
	var pagination = NewPagination[models.Joke]()

	j.db.Model([]models.Joke{}).Count(&totalRows)
	pagination.UpdateSettings(page, perPage)
	j.db.Scopes(paginate(pagination)).Find(&jokes)

	return pagination.PopulateData(totalRows, jokes)
}

// JokeExistInLastMonth check if same joke exist in database ,and it is newer then month
func (j *Joke) JokeExistInLastMonth(joke *models.Joke) bool {
	monthEarlier := time.Now().AddDate(0, -1, 0)

	r := j.db.
		Where("id = ? AND created_at > ?", joke.ID, monthEarlier.String()).Limit(1).Find(&joke)
	if r.Error != nil {
		log.Println(r.Error)
	}

	return r.RowsAffected > 0
}

func (j *Joke) GetFavourites(page, perPage int, userID uint) *Pagination[models.Joke] {
	var totalRows int64
	var jokes []models.Joke
	var pagination = NewPagination[models.Joke]()
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
