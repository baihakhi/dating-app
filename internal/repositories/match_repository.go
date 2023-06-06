package repositories

import (
	"fmt"

	"github.com/baihakhi/dating-app/internal/repositories/queries"
)

func (r *repository) CreateMatch(user1, user2 uint64) (int64, error) {
	var result int64
	if err := r.db.QueryRow(queries.CreateMatch, user1, user2).Scan(&result); err != nil {
		return 0, err
	}
	fmt.Println("XxX", user1, user2, result)

	return result, nil
}
