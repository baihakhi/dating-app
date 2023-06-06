package services

import "github.com/baihakhi/dating-app/internal/models"

func (s *service) CreateSwipe(username string, swiper, swiped uint64, is_liked bool) (int64, error) {
	swipeLeft, err := s.repositories.RedisUserGetSwipes(username)
	if err != nil {
		return 0, err
	}

	if err := s.repositories.RedisUserSetSwipes(username, swipeLeft-1, 0); err != nil {
		return 0, err
	}

	return s.repositories.CreateSwipe(swiper, swiped, is_liked)
}

func (s *service) GetSwipe(swiperID, userID uint64) (*models.Swipe, error) {
	return s.repositories.GetSwipe(swiperID, userID)
}

func (s *service) DeleteSwipe(userID uint64) error {
	return s.repositories.DeleteSwipe(userID)
}
