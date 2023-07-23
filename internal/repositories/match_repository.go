package repositories

import (
	"github.com/baihakhi/dating-app/internal/repositories/queries"
)

// CreateMatch creates a new match between two users in the database.
// It takes the IDs of user1 and user2, as well as the swipe ID that resulted in the match.
// It returns the ID of the created match or an error if it fails.
func (r *repository) CreateMatch(user1, user2, swipeID uint64) (int64, error) {
	var result int64
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	// Create the match record
	if err := tx.QueryRow(queries.CreateMatch, user1, user2).Scan(&result); err != nil {
		tx.Rollback()
		return 0, err
	}

	// Delete the swipe record associated with the match
	if _, err := r.db.Exec(queries.DeleteSwipe, swipeID); err != nil {
		tx.Rollback()
		return 0, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return result, nil
}
