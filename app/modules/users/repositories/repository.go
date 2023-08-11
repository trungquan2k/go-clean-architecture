package repository

import (
	"github.com/1rhino/clean_architecture/app/models"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CheckEmailExisting(email string) bool
	CreateUser(data *models.SignUpInput) (*models.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}
