package services

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
