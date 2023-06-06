package queries

const (
	CreateMatch = `
	INSERT INTO matches (user1_id, user2_id)
	VALUES ($1,$2)
	RETURNING match_id
	`
)
