package queries

const (
	CreateSwipe = `
	INSERT INTO swipes (swiper_id, swiped_id, is_liked)
	VALUES ($1,$2, $3)
	RETURNING swipe_id
	`
)
