package repositories

import (
	"github.com/baihakhi/dating-app/internal/repositories/queries"
)

func (r *repository) CreateMatch(user1, user2 uint64) (int64, error) {
	var result int64
	if err := r.db.QueryRow(queries.CreateMatch, user1, user2).Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}
