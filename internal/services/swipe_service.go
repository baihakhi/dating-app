package services

import (
	"strings"
	"time"

	"github.com/baihakhi/dating-app/internal/models"
	"github.com/baihakhi/dating-app/internal/models/enum"
	response "github.com/baihakhi/dating-app/internal/models/payload/responses"
)

// CreateSwipe creates a new swipe for the given user.
// It returns the ID of the created swipe or an error if it fails.
func (s *service) CreateSwipe(username string, swipe *models.Swipe, lastLogin *time.Time) (int64, error) {
	// Get the remaining swipe count for the user from Redis
	swipeLeft, err := s.repositories.RedisUserGetSwipes(username)
	if err != nil {
		if lastLogin == nil {
			// If the user's last login time is nil, it means they are not logged in,
			// and handle this case by returning a specific error message.
			if strings.Contains(err.Error(), "redis: nil") {
				err = response.ErrorBuilder(string(models.ACC), response.ULGN)
			}
			// Other error cases are returned as is.
			return 0, err
		} else {
			// If the user's last login time is not nil, because of
			// their Redis data is missing, set the swipe count by initial value.
			if err := s.redisUserSetSwipes(username); err != nil {
				return 0, err
			}
			swipeLeft = enum.InitialSwipes
		}
	}

	// Decrement the swipe count by 1 and update it in Redis
	if err := s.repositories.RedisUserSetSwipes(username, swipeLeft-1, 0); err != nil {
		return 0, err
	}

	// Create the swipe record in the database
	return s.repositories.CreateSwipe(swipe.Swiper, swipe.Swiped, swipe.IsLiked)
}

// GetSwipe retrieves a swipe record from the database based on the swiper ID and user ID.
func (s *service) GetSwipe(swiperID, userID uint64) (*models.Swipe, error) {
	return s.repositories.GetSwipe(swiperID, userID)
}

