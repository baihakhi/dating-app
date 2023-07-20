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
	Logout(data *models.User) (error)
	RemoveSwipeLimit(username string) error

	// Swipe Services
	CreateSwipe(username string, swiper, swiped uint64, is_liked bool) (int64, error)
	GetSwipe(swiperID, userID uint64) (*models.Swipe, error)
	DeleteSwipe(userID uint64) error

	// Match Services
	CreateMatch(user1, user2 uint64) (int64, error)
}
