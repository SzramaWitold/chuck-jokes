package controllers

import (
	"chuck-jokes/pkg/api/controllers/requests"
	"chuck-jokes/pkg/api/controllers/responses"
	"chuck-jokes/pkg/repositories"
	"chuck-jokes/pkg/token"
)

type ControllerWrapper struct {
	Category CategoryController
	Joke     JokeController
	User     UserController
}

func NewControllerWrapper(
	jwt *token.TokenHandler,
	request requests.RequestHandler,
	response responses.ResponseHandler,
	repository *repositories.Repository,
) *ControllerWrapper {
	return &ControllerWrapper{
		Category: NewCategory(request, response, repository),
		Joke:     NewJoke(request, response, repository),
		User:     NewUser(request, response, repository, jwt),
	}
}
