package services

import (
	"github.com/baihakhi/dating-app/internal/models"
	"github.com/baihakhi/dating-app/internal/repositories"
)

type service struct {
	repositories repositories.Repositories
}

func InitService(repo repositories.Repositories) Services {
	return &service{
		repositories: repo,
	}
}

type Services interface {
	// User Services
	CreateUser(data *models.User) (string, error)
	GetOneUserByUsername(username string) (*models.User, error)
	PatchUserVerified(userID uint64) error
	NextUser(userID uint64) (*models.User, error)
	Login(data *models.User) (string, error)
}
