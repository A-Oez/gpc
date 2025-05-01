CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE,
  age INT NOT NULL,
);

INSERT INTO users (username, email, age) VALUES
  ('alice', 'alice@example.com', 28),
  ('bob', 'bob@example.com', 34),
  ('charlie', 'charlie@example.com', 22);
