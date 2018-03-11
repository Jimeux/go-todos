-- Schema

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL
);

CREATE INDEX ON users (username, password);

CREATE TABLE IF NOT EXISTS todos (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users(id),
  title VARCHAR(255) NOT NULL,
  complete BOOLEAN NOT NULL DEFAULT FALSE,
  created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- Seed data

INSERT INTO users (username, password) VALUES ('jim', 'pass');

INSERT INTO todos (user_id, title) VALUES (1, 'Create lots of bugs');
INSERT INTO todos (user_id, title) VALUES (1, 'Fix lots of bugs');
