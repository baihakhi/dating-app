package repositories

import (
	"github.com/baihakhi/dating-app/internal/models"
	"github.com/baihakhi/dating-app/internal/repositories/queries"
)

// CreateSwipe creates a new swipe record in the database.
// It returns the ID of the created swipe record or an error if it fails.
func (r *repository) CreateSwipe(swiper, swiped uint64, isLiked bool) (int64, error) {
	var result int64
	err := r.db.QueryRow(queries.CreateSwipe, swiper, swiped, isLiked).Scan(&result)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// GetSwipe retrieves a swipe record from the database based on the swiper ID and user ID.
// It returns a pointer to the retrieved Swipe model or an error if it fails.
func (r *repository) GetSwipe(swiperID, userID uint64) (*models.Swipe, error) {
	result := new(models.Swipe)

	err := r.db.QueryRow(queries.GetSwipe, swiperID, userID).Scan(&result.ID, &result.Swiper, &result.Swiped, &result.IsLiked)
	if err != nil {
		return nil, err
	}

	return result, nil
}

