package gorm

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
	FindAll(request requests.FindCollection) *Pagination[models.Joke]
	JokeExistInLastMonth(joke *models.Joke) bool
	FindFavourites(request requests.FindCollection, userID uint) *Pagination[models.Joke]
	Create(joke *models.Joke) (*models.Joke, error)
}

// Joke base mapJoke repository.
type Joke struct {
	db *gorm.DB
}

// NewJoke create new mapJoke repository.
func NewJoke(db *gorm.DB) *Joke {
	return &Joke{
		db: db,
	}
}

// JokeOfTheDay get mapJoke of the day from specify day.
func (j *Joke) JokeOfTheDay(time string) *models.Joke {
	joke := gormModels.Joke{}

	if tx := j.db.Where("created_at >= ?", time).First(&joke); tx.Error != nil {
		log.Println(tx.Error)

		return nil
	}

	return mapJoke(&joke)
}

// Find get mapJoke of the day from specify day.
func (j *Joke) Find(jokeID uint) *models.Joke {
	joke := gormModels.Joke{}

	if txFind := j.db.First(&joke, jokeID); txFind.Error != nil {
		log.Println(txFind.Error)

		return nil
	}

	if txSave := j.db.Save(&joke); txSave.Error != nil {
		log.Println(txSave.Error)

		return nil
	}

	return mapJoke(&joke)
}

// FindAll get all jokes.
func (j *Joke) FindAll(request requests.FindCollection) *Pagination[models.Joke] {
	var totalRows int64

	var jokes []gormModels.Joke

	pagination := NewPagination[models.Joke]()

	j.db.Model([]gormModels.Joke{}).Count(&totalRows)
	pagination.UpdateSettings(request.Page, request.PerPage)
	j.db.Scopes(paginate(pagination)).Find(&jokes)

	baseJokes := make([]models.Joke, 0, len(jokes))

	for _, joke := range jokes {
		j := joke
		baseJoke := mapJoke(&j)
		baseJokes = append(baseJokes, *baseJoke)
	}

	return pagination.PopulateData(totalRows, baseJokes)
}

// JokeExistInLastMonth check if same mapJoke exist in database ,and it is in current month.
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
		jokes      []gormModels.Joke
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

	baseJokes := make([]models.Joke, 0, len(jokes))

	for _, joke := range jokes {
		j := joke
		baseJoke := mapJoke(&j)
		baseJokes = append(baseJokes, *baseJoke)
	}

	return pagination.PopulateData(totalRows, baseJokes)
}

func (j *Joke) Create(joke *models.Joke) (*models.Joke, error) {
	gormJoke := gormModels.Joke{
		Value:      joke.Value,
		ExternalID: joke.ExternalID,
	}

	gormJoke.CreatedAt = joke.CreatedAt
	gormJoke.UpdatedAt = joke.UpdatedAt

	if tx := j.db.Create(&gormJoke); tx.Error != nil {
		return nil, tx.Error
	}

	jsRepository := NewJokeStatistic(j.db)

	if err := jsRepository.Create(gormJoke.ID); err != nil {
		return nil, err
	}

	return mapJoke(&gormJoke), nil
}
