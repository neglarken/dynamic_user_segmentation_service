CREATE TABLE slugs
(
  id    BIGINT PRIMARY KEY,
  title TEXT    NOT NULL
);

CREATE TABLE users
(
  id BIGINT PRIMARY KEY 
);

CREATE TABLE slugs_users
(
  user_id INTEGER REFERENCES users (id),
  slug_id INTEGER REFERENCES slugs (id)
);     