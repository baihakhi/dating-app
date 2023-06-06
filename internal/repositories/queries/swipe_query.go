package queries

const (
	CreateSwipe = `
	INSERT INTO swipes (swiper_id, swiped_id, is_liked)
	VALUES ($1,$2, $3)
	RETURNING swipe_id
	`

	DeleteSwipe = `
	DELETE from swipes
	WHERE
		swipes.swipe_id = $1
	`

	GetSwipe = `
	SELECT
		swipe_id,
		swiper_id,
		swiped_id,
		is_liked
	FROM swipes
	WHERE 
		swiper_id = $1
	AND
		swiped_id = $2
	LIMIT 1
	`
)
