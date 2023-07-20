package repositories

import (
	"encoding/json"
	"time"

	"github.com/baihakhi/dating-app/internal/models"
	"github.com/baihakhi/dating-app/internal/repositories/queries"
)

// CreateUser creates a new user in the database with the provided data.
func (r *repository) CreateUser(data *models.User) (string, error) {
	var username string
	if err := r.db.QueryRow(queries.CreateUser,
		data.Username,
		data.Email,
		data.Password,
		data.Fullname,
		data.Gender,
		data.Preference,
		data.City,
		data.Interest).
		Scan(&username); err != nil {
		return "", err
	}

	return username, nil
}

// GetOneUsersByUsername retrieves a single user from the database based on the provided username.
func (r *repository) GetOneUsersByUsername(username string) (*models.User, error) {
	var result models.User

	if err := r.db.QueryRow(queries.GetOneUserByUsername, username).
		Scan(
			&result.ID,
			&result.Username,
			&result.Email,
			&result.Fullname,
			&result.Gender,
			&result.Preference,
			&result.City,
			&result.Interest,
			&result.IsVerified,
			&result.LastLogin,
			&result.CreatedAt,
			&result.UpdatedAt,
		); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPasswordByUsername retrieves the password for a given username from the database.
func (r *repository) GetPasswordByUsername(username string) (string, error) {
	var result string

	if err := r.db.QueryRow(queries.GetPasswordByUsername, username).
		Scan(&result); err != nil {
		return "", err
	}
	return result, nil
}

// PatchUserVerified updates the verification status of a user in the database.
func (r *repository) PatchUserVerified(userID uint64) error {
	_, err := r.db.Exec(queries.PatchUserVerified, userID)
	return err
}

// PatchUserLogin updates the last_login value of a user in the database.
func (r *repository) PatchUserLogin(userID uint64) error {
	_, err := r.db.Exec(queries.PatchUserLogin, userID)
	return err
}

// PatchUserLogout removes the last_login value of a user in the database.
func (r *repository) PatchUserLogout(userID uint64) error {
	_, err := r.db.Exec(queries.PatchUserLogout, userID)
	return err
}

// NextUser retrieves the next user from the database based on the provided userID.
func (r *repository) NextUser(userID uint64) (*models.User, error) {
	var result models.User

	if err := r.db.QueryRow(queries.NextUser, userID).
		Scan(
			&result.ID,
			&result.Username,
			&result.Email,
			&result.Fullname,
			&result.Gender,
			&result.Preference,
			&result.City,
			&result.Interest,
		); err != nil {
		return nil, err
	}
	return &result, nil
}

// RedisUserSwipes set max swipes per day in redis.
func (r *repository) RedisUserSetSwipes(username string, swipes int, t time.Duration) error {
	return r.redis.HSet(username, models.USwipes, swipes, t*time.Hour)
}

// RedisUserSwipes remove max swipes per day in redis.
func (r *repository) RedisUserRemoveLImit(username string) error {
	return r.redis.HDel(username, models.USwipes)
}

// RedisUserGetSwipes get swipes left from redis.
func (r *repository) RedisUserGetSwipes(username string) (int, error) {
	var swipes int
	result, err := r.redis.HGet(username, models.USwipes)
	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(result, &swipes); err != nil {
		return 0, err
	}

	return swipes, nil
}
