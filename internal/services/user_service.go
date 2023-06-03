package services

import (
	"strings"

	"github.com/baihakhi/dating-app/internal/middleware"
	"github.com/baihakhi/dating-app/internal/models"
	"github.com/baihakhi/dating-app/internal/models/enum"
	hash "github.com/baihakhi/dating-app/internal/utils/bycript"
)

func (s *service) CreateUser(data *models.User) (string, error) {
	hashedPass, err := hash.Encrypt(data.Password)
	if err != nil {
		return "", err
	}
	data.Password = hashedPass
	return s.repositories.CreateUser(data)
}

func (s *service) GetOneUserByUsername(username string) (*models.User, error) {
	return s.repositories.GetOneUsersByUsername(username)
}

func (s *service) PatchUserVerified(userID uint64) error {
	return s.repositories.PatchUserVerified(userID)
}

func (s *service) NextUser(userID uint64) (*models.User, error) {
	return s.repositories.NextUser(userID)
}

func (s *service) Login(data *models.User) (string, error) {
	pass, err := s.repositories.GetPasswordByUsername(strings.ToLower(data.Username))
	if err != nil {
		return "", err
	}

	if err := hash.VerifyPassword(pass, data.Password); err != nil {
		return "", err
	}

	acc, err := s.repositories.GetOneUsersByUsername(strings.ToLower(data.Username))
	if err != nil {
		return "", err
	}

	if err := s.repositories.PatchUserLogin(acc.ID); err != nil {
		return "", err
	}

	if err := s.repositories.RedisUserSwipes(acc.Username, enum.InitialSwipes); err != nil {
		return "", err
	}

	return middleware.CreateToken(*acc)
}
