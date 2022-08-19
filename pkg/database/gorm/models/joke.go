package models

import (
	"chuck-jokes/models"
	"log"

	"gorm.io/gorm"
)

// Joke gorm model needed for migration
type Joke struct {
	gorm.Model
	Value      string `faker:"sentence"`
	ExternalID string `gorm:"unique" faker:"unique"`
}

//func ChangeToGormModel(externalModel *models.Joke) *Joke {
//	model := gorm.Model{
//		ID:        externalModel.ID,
//		CreatedAt: externalModel.CreatedAt,
//		UpdatedAt: externalModel.UpdatedAt,
//		DeletedAt: gorm.DeletedAt(externalModel.DeletedAt),
//	}
//	return &Joke{
//		Model:      model,
//		Value:      externalModel.Value,
//		ExternalID: externalModel.ExternalID,
//	}
//}

//func (j *Joke) ChangeToBaseModel() *models.Joke {
//	return &models.Joke{
//		ID:         j.ID,
//		CreatedAt:  j.CreatedAt,
//		DeletedAt:  sql.NullTime(j.DeletedAt),
//		UpdatedAt:  j.UpdatedAt,
//		Value:      j.Value,
//		ExternalID: j.ExternalID,
//	}
//}

// Create save joke to database
func Create(db *gorm.DB, joke *models.Joke) *models.Joke {
	tx := db.Create(&joke)
	if tx.Error != nil {
		log.Println(tx.Error)
	}

	return joke
}

// GetID from Joke
func (j *Joke) GetID() uint {
	return j.ID
}

// ExternalJoke for call from external api
type ExternalJoke struct {
	Value string `faker:"sentence"`
	ID    string `gorm:"primaryKey" faker:"unique"`
}

//// ChangeToGormJoke transfer data from ExternalJoke to DBJoke
//func ChangeToGormJoke(joke *ExternalJoke) Joke {
//	return Joke{
//		Value:      joke.Value,
//		ExternalID: joke.ID,
//	}
//}
