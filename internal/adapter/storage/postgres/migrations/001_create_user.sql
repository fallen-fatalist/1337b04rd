CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(70) NOT NULL,
    avatar VARCHAR(70) NOT NULL,
    session_id UUID DEFAULT uuid_generate_v4(),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    expires_at TIMESTAMPTZ
);