package services

import (
	"strings"

	"github.com/baihakhi/dating-app/internal/middleware"
	"github.com/baihakhi/dating-app/internal/models"
	"github.com/baihakhi/dating-app/internal/models/enum"
	hash "github.com/baihakhi/dating-app/internal/utils/bycript"
)

// CreateUser creates a new user by encrypting the password, updating it in the user model,
// It returns the username of the created user if successful; otherwise, it returns an error.
func (s *service) CreateUser(data *models.User) (string, error) {
	// Encrypt the user's password before storing it in the database.
	hashedPass, err := hash.Encrypt(data.Password)
	if err != nil {
		return "", err
	}
	data.Password = hashedPass

	// Call the CreateUser method of the repository layer to store the user in the database.
	return s.repositories.CreateUser(data)
}

// GetOneUserByUsername retrieves a user from the database based on the provided username.
// It returns the user model if found; otherwise, it returns an error.
func (s *service) GetOneUserByUsername(username string) (*models.User, error) {
	return s.repositories.GetOneUsersByUsername(username)
}

// PatchUserVerified updates the user's verified status in the database.
func (s *service) PatchUserVerified(userID uint64) error {
	return s.repositories.PatchUserVerified(userID)
}

// RemoveSwipeLimit removes the swipe limit for a user in Redis, allowing unlimited swipes.
func (s *service) RemoveSwipeLimit(username string) error {
	return s.repositories.RedisUserRemoveLImit(username)
}

// NextUser fetches the next user to be recommended to the user with the given userID.
func (s *service) NextUser(userID uint64) (*models.User, error) {
	return s.repositories.NextUser(userID)
}

// Login authenticates the user by verifying the provided username and password.
// verifies the provided password against the stored password hash,
// updates the last login time of the user, and creates a JWT token for the user.
// If the login is successful, it returns the JWT token; otherwise, it returns an error.
func (s *service) Login(data *models.User) (string, error) {
	// Get the password hash from the database based on the provided username.
	pass, err := s.repositories.GetPasswordByUsername(strings.ToLower(data.Username))
	if err != nil {
		return "", err
	}

	// Verify the provided password against the stored password hash using a cryptographic hash function.
	if err := hash.VerifyPassword(pass, data.Password); err != nil {
		return "", err
	}

	// Retrieve the user account from the database based on the username.
	acc, err := s.repositories.GetOneUsersByUsername(strings.ToLower(data.Username))
	if err != nil {
		return "", err
	}

	// Update the last login time of the user in the database.
	if err := s.repositories.PatchUserLogin(acc.ID); err != nil {
		return "", err
	}

	// If the user account is not verified, set the swipe count for the user in Redis (assumes new user or missing Redis data).
	if !acc.IsVerified {
		if err := s.redisUserSetSwipes(acc.Username); err != nil {
			return "", err
		}
	}

	// Create a JWT token for the user and return it.
	return middleware.CreateToken(*acc)
}

// redisUserSetSwipes sets the initial swipe count for a new user in Redis.
// It returns an error if there is an issue while setting the swipe count in Redis.
func (s *service) redisUserSetSwipes(username string) error {
	return s.repositories.RedisUserSetSwipes(username, enum.InitialSwipes, 24)
}

// Logout updates the user's logout status by updating the last logout time in the database.
func (s *service) Logout(data *models.User) error {
	// Retrieve the user account from the database based on the username.
	acc, err := s.repositories.GetOneUsersByUsername(strings.ToLower(data.Username))
	if err != nil {
		return err
	}

	// Update the last logout time of the user in the database.
	if err := s.repositories.PatchUserLogout(acc.ID); err != nil {
		return err
	}

	return nil
}
