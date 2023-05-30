package repositories

import (
	"database/sql"

	"github.com/baihakhi/dating-app/internal/models"
)

type repository struct {
	db *sql.DB
}

func InitRepository(db *sql.DB) Repositories {
	return &repository{
		db: db,
	}
}

type Repositories interface {
	// User Repository
	CreateUser(data *models.User) (string, error)
	GetOneUsersByUsername(username string) (*models.User, error)
	GetPasswordByUsername(username string) (string, error)
	PatchUserVerified(userID uint64) error
	NextUser(userID uint64) (*models.User, error)
}
