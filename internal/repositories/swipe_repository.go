package repositories

import (
	"github.com/baihakhi/dating-app/internal/models"
	"github.com/baihakhi/dating-app/internal/repositories/queries"
)

func (r *repository) CreateSwipe(swiper, swiped uint64, is_liked bool) (int64, error) {
	var result int64
	err := r.db.QueryRow(queries.CreateSwipe, swiper, swiped, is_liked).Scan(&result)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (r *repository) GetSwipe(swiperID, userID uint64) (*models.Swipe, error) {
	result := new(models.Swipe)

	err := r.db.QueryRow(queries.GetSwipe, swiperID, userID).Scan(&result.ID, &result.Swiper, &result.Swiped, &result.IsLiked)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) DeleteSwipe(userID uint64) error {
	_, err := r.db.Exec(queries.DeleteSwipe, userID)

	return err
}
