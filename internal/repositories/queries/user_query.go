package queries

const (
	CreateUser = `
	INSERT INTO users (username, email, password, full_name, gender, preference, city, interest)
	VALUES ($1, $2, $3, $4, $5, $6, #7, $8)
	RETURNING username
	`

	GetOneUserByUsername = `
	SELECT 
		user_id,
		username, 
		email, 
		full_name, 
		gender, 
		preference, 
		city, 
		interest,
		is_verified,
		created_at,
		updated_at
	FROM users
	WHERE username = $1
	LIMIT 1
	`

	GetPasswordByUsername = `
	SELECT
		password
	FROM users
	WHERE
		username = $1
	LIMIT 1
	`
	PatchUserVerified = `
	UPDATE users
	SET
		is_verified = true
	WHERE 
		user_id = $1
	`

	NextUser = `
	WITH current_user AS (
		SELECT *
		FROM users
		WHERE 
			user_id = $1
	), current_user_interest AS (
		SELECT unnest(string_to_array(interest, ', ')) AS interest
		FROM current_user
	)
	SELECT 
		u.username, 
		u.email, 
		u.full_name, 
		u.gender, 
		u.preference,
		u.city, 
		u.interest
	FROM users u
	JOIN current_user_interest cui ON u.interest ILIKE '%' || cui.interest || '%'
	WHERE u.user_id <> $1
		AND user_id NOT IN (
			SELECT swiped_id
			FROM swipes
			WHERE swiper_id = $1
			UNION
			SELECT user_id
			FROM matches
			WHERE user1_id = $1 OR user2_id = $1
		)
		OR cui.interest IS NULL
	ORDER BY RANDOM()
	LIMIT 1;
	`
)
