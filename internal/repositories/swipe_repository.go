package repositories

import "github.com/baihakhi/dating-app/internal/repositories/queries"

func (r *repository) CreateSwipe(swiper, swiped uint64, is_liked bool) (int64, error) {
	var result int64
	err := r.db.QueryRow(queries.CreateSwipe, swiper, swiped, is_liked).Scan(&result)
	if err != nil {
		return 0, err
	}

	return result, nil
}
