package services

// CreateMatch creates a new match for the given users.
func (s *service) CreateMatch(user1, user2, swipeID uint64) (int64, error) {
	return s.repositories.CreateMatch(user1, user2, swipeID)
}
