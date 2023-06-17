package services

import "github.com/baihakhi/dating-app/internal/models"

// CreateSwipe creates a new swipe for the given username.
// It takes the username of the swiper, the IDs of the swiper and swiped users, and a boolean indicating whether the swipe is a like or dislike.
// It returns the ID of the created swipe or an error if it fails.
func (s *service) CreateSwipe(username string, swiper, swiped uint64, isLiked bool) (int64, error) {
	// Get the remaining swipe count for the user from Redis
	swipeLeft, err := s.repositories.RedisUserGetSwipes(username)
	if err != nil {
		return 0, err
	}

	// Decrement the swipe count by 1 and update it in Redis
	if err := s.repositories.RedisUserSetSwipes(username, swipeLeft-1, 0); err != nil {
		return 0, err
	}

	// Create the swipe record in the database
	return s.repositories.CreateSwipe(swiper, swiped, isLiked)
}

// GetSwipe retrieves a swipe record from the database based on the swiper ID and user ID.
// It returns a pointer to the retrieved Swipe model or an error if it fails.
func (s *service) GetSwipe(swiperID, userID uint64) (*models.Swipe, error) {
	return s.repositories.GetSwipe(swiperID, userID)
}

