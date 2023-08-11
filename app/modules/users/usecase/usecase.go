package usecase

import (
	"github.com/1rhino/clean_architecture/app/models"
	users "github.com/1rhino/clean_architecture/app/modules/users/repositories"
	"github.com/gofiber/fiber/v2"
)

type UseCase interface {
	SignUpUser(ctx *fiber.Ctx, payload *models.SignUpInput) (*models.UserResponse, error)
}

type UserUseCase struct {
	userRepo users.UserRepoInterface
}

func NewUserUseCase(userRepo users.UserRepoInterface) UseCase {
	return &UserUseCase{userRepo: userRepo}
}
