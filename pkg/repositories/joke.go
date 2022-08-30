package repositories

import (
	"log"
	"time"

	"chuck-jokes/models"
	"chuck-jokes/pkg/api/controllers/requests"
	gormModels "chuck-jokes/pkg/database/gorm/models"

	"gorm.io/gorm"
)

type JokeRepository interface {
	JokeOfTheDay(time string) *models.Joke
	Find(jokeID uint) *models.Joke
	GetStatistic(jokeID uint) (*models.Joke, uint)
	FindAll(request requests.FindCollection) *Pagination[models.Joke]
	JokeExistInLastMonth(joke *models.Joke) bool
	FindFavourites(request requests.FindCollection, userID uint) *Pagination[models.Joke]
}

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

// JokeOfTheDay get joke of the day from specify day
func (j *Joke) JokeOfTheDay(time string) *models.Joke {
	joke := models.Joke{}

	if tx := j.db.Where("created_at >= ?", time).First(&joke); tx.Error != nil {
		log.Println(tx.Error)

		return nil
	}

	return &joke
}

// Find get joke of the day from specify day
func (j *Joke) Find(jokeID uint) *models.Joke {
	joke := models.Joke{}

	if txFind := j.db.First(&joke, jokeID); txFind.Error != nil {
		log.Println(txFind.Error)

		return nil
	}

	joke.Shows++

	if txSave := j.db.Save(&joke); txSave.Error != nil {
		log.Println(txSave.Error)

		return nil
	}

	return &joke
}

// GetStatistic get statistic for joke
func (j *Joke) GetStatistic(jokeID uint) (*models.Joke, uint) {
	joke := models.Joke{}

	if txFind := j.db.First(&joke, jokeID); txFind.Error != nil {
		log.Println(txFind.Error)

		return nil, 0
	}

	favourites := j.db.Model(&joke).Association("Users").Count()

	return &joke, uint(favourites)
}

// FindAll get all jokes
func (j *Joke) FindAll(request requests.FindCollection) *Pagination[models.Joke] {
	var totalRows int64

	var jokes []models.Joke

	pagination := NewPagination[models.Joke]()

	j.db.Model([]models.Joke{}).Count(&totalRows)
	pagination.UpdateSettings(request.Page, request.PerPage)
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

func (j *Joke) FindFavourites(request requests.FindCollection, userID uint) *Pagination[models.Joke] {
	var (
		totalRows  int64
		jokes      []models.Joke
		pagination = NewPagination[models.Joke]()
		user       = gormModels.User{}
	)

	if tx := j.db.First(&user, userID); tx.Error != nil {
		log.Println(tx.Error)

		return nil
	}

	totalRows = j.db.Model(&user).Association("Favourites").Count()

	pagination.UpdateSettings(request.Page, request.PerPage)

	dbError := j.db.Model(&user).Scopes(paginate(pagination)).Association("Favourites").Find(&jokes)

	if dbError != nil {
		log.Println(dbError)

		return nil
	}

	return pagination.PopulateData(totalRows, jokes)
}
