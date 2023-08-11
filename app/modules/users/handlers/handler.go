package handlers

import (
	user "github.com/1rhino/clean_architecture/app/modules/users/usecase"
)

type UserHandlers struct {
	userUseCase user.UseCase
}

func NewUserHandlers(userUseCase user.UseCase) *UserHandlers {
	return &UserHandlers{userUseCase: userUseCase}
}
