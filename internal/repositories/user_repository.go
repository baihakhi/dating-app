package repositories

import (
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

// NextUser retrieves the next user from the database based on the provided userID.
func (r *repository) NextUser(userID uint64) (*models.User, error) {
	var result models.User

	if err := r.db.QueryRow(queries.NextUser, userID).
		Scan(
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
