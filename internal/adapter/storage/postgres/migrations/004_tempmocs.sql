-- Enable the extension (already in your schema)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- 👤 Users
INSERT INTO users (name, avatar, expires_at) VALUES
('Rick', 'https://rickandmortyapi.com/api/character/avatar/1.jpeg', NOW() + INTERVAL '7 days'),
('Morty', 'https://rickandmortyapi.com/api/character/avatar/2.jpeg', NOW() + INTERVAL '7 days'),
('Summer', 'https://rickandmortyapi.com/api/character/avatar/3.jpeg', NOW() + INTERVAL '7 days'),
('Beth', 'https://rickandmortyapi.com/api/character/avatar/4.jpeg', NOW() + INTERVAL '7 days'),
('Jerry', 'https://rickandmortyapi.com/api/character/avatar/5.jpeg', NOW() + INTERVAL '7 days');

-- 🧵 Posts
INSERT INTO posts (user_id, title, content, image) VALUES
(1, 'Welcome Hackers', 'First post on the board!', 'https://s3.example.com/bucket1/welcome.jpg'),
(2, '1337 Tips', 'Remember to hide your IP.', NULL),
(3, 'Ask Anything', 'Thread for random questions', 'https://s3.example.com/bucket2/questions.png');

-- 💬 Comments
INSERT INTO comments (user_id, post_id, content) VALUES
(2, 1, 'Nice board!'),
(3, 1, 'Can I post images?'),
(1, 3, 'Feel free to ask anything.');
