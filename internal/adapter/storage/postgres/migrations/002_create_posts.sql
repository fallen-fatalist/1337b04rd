CREATE TABLE posts (
    post_id BIGSERIAL PRIMARY KEY,
    user_id TEXT NOT NULL, 
    title VARCHAR(20) NOT NULL,
    content TEXT NOT NULL,
    image TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(user_id) 
);

--     expires_at TIMESTAMPTZ,
--     active BOOLEAN 
