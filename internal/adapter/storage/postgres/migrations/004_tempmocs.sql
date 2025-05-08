-- Enable the extension (already in your schema)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- 👤 Users
INSERT INTO users (user_id, name, avatar, expires_at) VALUES
(uuid_generate_v4(), 'Rick', 'https://rickandmortyapi.com/api/character/avatar/1.jpeg', NOW() + INTERVAL '7 days'),
(uuid_generate_v4(), 'Morty', 'https://rickandmortyapi.com/api/character/avatar/2.jpeg', NOW() + INTERVAL '7 days'),
(uuid_generate_v4(), 'Summer', 'https://rickandmortyapi.com/api/character/avatar/3.jpeg', NOW() + INTERVAL '7 days'),
(uuid_generate_v4(), 'Beth', 'https://rickandmortyapi.com/api/character/avatar/4.jpeg', NOW() + INTERVAL '7 days'),
(uuid_generate_v4(), 'Jerry', 'https://rickandmortyapi.com/api/character/avatar/5.jpeg', NOW() + INTERVAL '7 days');

-- For referencing users later, create temp table or CTE to get their UUIDs
WITH u AS (
  SELECT user_id, name FROM users
)

-- 🧵 Posts (replace with real UUIDs after user insert)
INSERT INTO posts (user_id, title, content, image) VALUES
((SELECT user_id FROM users WHERE name = 'Rick'), 'Welcome Hackers', 'First post on the board!', 'https://s3.example.com/bucket1/welcome.jpg'),
((SELECT user_id FROM users WHERE name = 'Morty'), '1337 Tips', 'Remember to hide your IP.', NULL),
((SELECT user_id FROM users WHERE name = 'Summer'), 'Ask Anything', 'Thread for random questions', 'https://s3.example.com/bucket2/questions.png');

-- 💬 Comments (post_id references above)
INSERT INTO comments (user_id, post_id, content) VALUES
((SELECT user_id FROM users WHERE name = 'Morty'), 1, 'Nice board!'),
((SELECT user_id FROM users WHERE name = 'Summer'), 1, 'Can I post images?'),
((SELECT user_id FROM users WHERE name = 'Rick'), 3, 'Feel free to ask anything.');

--  GOVNO MOKI!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!