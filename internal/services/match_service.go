package services

func (s *service) CreateMatch(user1, user2 uint64) (int64, error) {
	return s.repositories.CreateMatch(user1, user2)
}
