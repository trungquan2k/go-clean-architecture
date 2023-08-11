package usecase

import (
	"errors"

	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gofiber/fiber/v2"
)

func (u UserUseCase) SignUpUser(ctx *fiber.Ctx, payload *models.SignUpInput) (*models.UserResponse, error) {
	if payload.Password != payload.PasswordConfirm {
		return nil, errors.New("passwords do not match")
	}

	// check existing email
	existing := u.userRepo.CheckEmailExisting(payload.Email)

	if existing == true {
		return nil, errors.New("email existing, please choose another email.")
	}

	createdUser, err := u.userRepo.CreateUser(payload)

	if err != nil {
		return nil, err
	}

	return models.FilterUserRecord(createdUser), nil
}
