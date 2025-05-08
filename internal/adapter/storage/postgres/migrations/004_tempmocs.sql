-- -- Insert mock users
-- INSERT INTO users (user_id, name, avatar, expires_at) VALUES
-- ('user-123', 'Rick Sanchez', 'https://rickandmortyapi.com/api/character/avatar/1.jpeg', NOW() + INTERVAL '7 days'),
-- ('user-456', 'Morty Smith', 'https://rickandmortyapi.com/api/character/avatar/2.jpeg', NOW() + INTERVAL '7 days');

-- -- Insert mock posts
-- INSERT INTO posts (user_name, user_avatar, title, content, image) VALUES
-- ('Rick Sanchez', 'https://rickandmortyapi.com/api/character/avatar/1.jpeg', 'First Post', 'Wubba Lubba Dub Dub!', NULL),
-- ('Morty Smith', 'https://rickandmortyapi.com/api/character/avatar/2.jpeg', 'Morty’s Thoughts', 'Oh geez, Rick...', NULL);

-- -- Insert mock comments
-- INSERT INTO comments (user_name, user_avatar, post_id, content) VALUES
-- ('Morty Smith', 'https://rickandmortyapi.com/api/character/avatar/2.jpeg', 1, 'Nice post, Rick!'),
-- ('Rick Sanchez', 'https://rickandmortyapi.com/api/character/avatar/1.jpeg', 2, 'Stay out of trouble, Morty.');
