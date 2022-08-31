package gorm

import (
	"time"

	"chuck-jokes/models"
	gormModel "chuck-jokes/pkg/database/gorm/models"
)

func mapJoke(joke *gormModel.Joke) *models.Joke {
	j := models.Joke{
		ID:         joke.ID,
		CreatedAt:  joke.CreatedAt,
		UpdatedAt:  joke.UpdatedAt,
		Value:      joke.Value,
		ExternalID: joke.ExternalID,
	}

	for _, u := range joke.Users {
		j.Users = append(j.Users, *mapUser(&u))
	}

	return &j
}

func mapJokeStatistic(statistic gormModel.JokeStatistic) *models.JokeStatistic {
	js := models.JokeStatistic{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Shows:     0,
		JokeID:    0,
		Joke:      models.Joke{},
	}

	js.Joke = *mapJoke(&statistic.Joke)

	return &js
}

func mapUser(user *gormModel.User) *models.User {
	u := models.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Username:  user.Username,
		Password:  user.Password,
	}

	for _, j := range user.Favourites {
		u.Favourites = append(u.Favourites, *mapJoke(&j))
	}

	for _, c := range user.Categories {
		u.Categories = append(u.Categories, *mapCategory(&c))
	}

	return &u
}

func mapCategory(category *gormModel.Category) *models.Category {
	c := models.Category{
		ID:        category.ID,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		Access:    category.Access,
		Name:      category.Name,
		Jokes:     nil,
		UserID:    category.UserID,
		User:      models.User{},
	}

	c.User = *mapUser(&category.User)

	for _, j := range category.Jokes {
		c.Jokes = append(c.Jokes, *mapJoke(&j))
	}

	return &c
}
