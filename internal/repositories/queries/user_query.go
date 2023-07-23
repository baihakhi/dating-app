package queries

const (
	CreateUser = `
	INSERT INTO users (username, email, password, full_name, gender, preference, city, interests)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
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
		interests,
		is_verified,
		last_login,
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

	PatchUserLogin = `
	UPDATE users
	SET
		last_login = now()
	WHERE 
		user_id = $1
	`

	PatchUserLogout = `
	UPDATE users
	SET
		last_login = NULL
	WHERE 
		user_id = $1
	`
	// NextUser selects the next user to recommend based on certain criteria.
	// It utilizes common table expressions (CTEs) to fetch the current user's details and interests.
	// The main query then joins these CTEs with the users table to find a random user who has at least
	// one common interest with the current user. The query also includes conditions to exclude users who
	// have already been swiped or matched with the current user.
	NextUser = `
	WITH curent_user AS (
		SELECT *
		FROM users
		WHERE 
			user_id = $1
	), curent_user_interests AS (
		SELECT unnest(string_to_array(interests, ', ')) AS interests
		FROM curent_user
	)
	SELECT 
		u.user_id,
		u.username, 
		u.email, 
		u.full_name, 
		u.gender, 
		u.preference,
		u.city, 
		u.interests
	FROM users u
	LEFT JOIN curent_user_interests cui ON u.interests ILIKE '%' || cui.interests || '%'
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
		OR cui.interests IS NULL
	ORDER BY RANDOM()
	LIMIT 1;
	`
)
