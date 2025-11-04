CREATE TABLE users (
                       id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                       full_name VARCHAR(255),
                       handle VARCHAR(150) NOT NULL UNIQUE,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(60) NOT NULL,
                       birth_date DATE,
                       avatar_url TEXT,
                       joined_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
                       is_deleted BOOLEAN DEFAULT FALSE,
                       deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE following_relationships (
                                         follower_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    /* changed followed to follower for better distinguish */
                                         following_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                                         followed_since TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                                         PRIMARY KEY (follower_id, following_id),
                                         CHECK (follower_id != following_id)
);

CREATE TABLE Posts (
                       id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                       user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                       imag_url TEXT,
                       caption VARCHAR(140),
                       created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_fr_following ON following_relationships(following_id);
CREATE INDEX idx_posts_user ON posts(user_id);
CREATE INDEX idx_posts_created_at ON posts(created_at);
