-- create user table
CREATE TABLE IF NOT EXISTS users (
    user_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    gender VARCHAR(50) DEFAULT 'not_specified'::character varying CHECK(gender IN ('male', 'female', 'not_specified')),
    preference VARCHAR(50) DEFAULT 'not_specified'::character varying CHECK(preference IN ('male', 'female', 'both', 'not_specified')),
    city VARCHAR(50),
    interests TEXT,
    is_verified BOOLEAN NOT NULL DEFAULT false,

    last_login TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- create swipes table
CREATE TABLE IF NOT EXISTS swipes (
    swipe_id BIGSERIAL PRIMARY KEY,
    swiper_id INT NOT NULL,
    swiped_id INT NOT NULL,
    is_liked BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (swiper_id) REFERENCES users(user_id) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (swiped_id) REFERENCES users(user_id) ON DELETE SET NULL ON UPDATE CASCADE
);

-- create matches table
CREATE TABLE IF NOT EXISTS matches (
    match_id BIGSERIAL PRIMARY KEY,
    user1_id INT NOT NULL,
    user2_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (user1_id) REFERENCES users(user_id) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (user2_id) REFERENCES users(user_id) ON DELETE SET NULL ON UPDATE CASCADE
);

-- adding users records to the table with password: `admin`
INSERT INTO users (username, email, password, full_name, gender, preference, city, interests)
VALUES 
  ('user1', 'user1@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'John Doe', 'male', 'female', 'New York', 'Photography, Hiking'),
  ('user2', 'user2@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Jane Smith', 'female', 'male', 'Los Angeles', 'Cooking, Yoga'),
  ('user3', 'user3@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Alex Johnson', 'male', 'both', 'Chicago', 'Music, Reading'),
  ('user4', 'user4@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Emily Davis', 'female', 'male', 'San Francisco', 'Art, Travel'),
  ('user5', 'user5@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Michael Wilson', 'male', 'not_specified', 'Seattle', 'Sports, Movies'),
  ('user6', 'user6@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Sarah Johnson', 'female', 'both', 'London', 'Dancing, Fashion'),
  ('user7', 'user7@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Robert Smith', 'male', 'female', 'Paris', 'Sports, Cooking'),
  ('user8', 'user8@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Sophia Davis', 'female', 'male', 'Sydney', 'Travel, Photography'),
  ('user9', 'user9@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Daniel Wilson', 'male', 'not_specified', 'Toronto', 'Gaming, Technology'),
  ('user10', 'user10@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Olivia Brown', 'female', 'female', 'Berlin', 'Reading, Music'),
  ('user11', 'user11@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'James Miller', 'male', 'both', 'Tokyo', 'Art, Movies'),
  ('user12', 'user12@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Ava Anderson', 'female', 'not_specified', 'Barcelona', 'Yoga, Cooking'),
  ('user13', 'user13@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'William Taylor', 'male', 'male', 'Rome', 'Hiking, History'),
  ('user14', 'user14@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Mia Thomas', 'female', 'female', 'Cape Town', 'Fashion, Travel'),
  ('user15', 'user15@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Benjamin Wilson', 'male', 'both', 'Dubai', 'Sports, Technology'),
  ('user16', 'user16@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Emma Wilson', 'female', 'male', 'New Delhi', 'Reading, Music'),
  ('user17', 'user17@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Alexander Lee', 'male', 'not_specified', 'Mumbai', 'Photography, Travel'),
  ('user18', 'user18@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Grace Thompson', 'female', 'both', 'Singapore', 'Dancing, Movies'),
  ('user19', 'user19@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Ethan Martin', 'male', 'female', 'Kuala Lumpur', 'Sports, Cooking'),
  ('user20', 'user20@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Chloe Lewis', 'female', 'male', 'Bangkok', 'Fashion, Yoga'),
  ('user21', 'user21@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Jacob Clark', 'male', 'both', 'Seoul', 'Music, Travel'),
  ('user22', 'user22@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Lily Turner', 'female', 'not_specified', 'Beijing', 'Art, Reading'),
  ('user23', 'user23@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Noah Rodriguez', 'male', 'male', 'Moscow', 'Gaming, Technology'),
  ('user24', 'user24@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Aria Hall', 'female', 'female', 'Mexico City', 'Cooking, Travel'),
  ('user25', 'user25@example.com', '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G', 'Matthew Young', 'male', 'both', 'Sao Paulo', 'Hiking, Movies');

-- create function for automated updated_at on update records
CREATE  FUNCTION update_timestamp_func()
RETURNS TRIGGER
LANGUAGE plpgsql AS
'
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
';

DO $$
DECLARE
    t text;
BEGIN
    FOR t IN
        SELECT table_name FROM information_schema.columns WHERE column_name = 'updated_at'
    LOOP
        EXECUTE format('CREATE TRIGGER trigger_update_timestamp
                    BEFORE UPDATE ON %I
                    FOR EACH ROW EXECUTE PROCEDURE update_timestamp_func()', t,t);
    END loop;
END;
$$ language 'plpgsql';