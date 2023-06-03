package repositories

import (
	"database/sql"

	driver "github.com/baihakhi/dating-app/internal/databases"
	"github.com/baihakhi/dating-app/internal/models"
)

type repository struct {
	db    *sql.DB
	redis driver.RedisClient
}

func InitRepository(db *sql.DB, redis driver.RedisClient) Repositories {
	return &repository{
		db:    db,
		redis: redis,
	}
}

type Repositories interface {
	// User Repository
	CreateUser(data *models.User) (string, error)
	GetOneUsersByUsername(username string) (*models.User, error)
	GetPasswordByUsername(username string) (string, error)
	PatchUserVerified(userID uint64) error
	PatchUserLogin(userID uint64) error
	NextUser(userID uint64) (*models.User, error)
	RedisUserSwipes(username string, swipes int) error
}
