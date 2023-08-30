CREATE TABLE IF NOT EXISTS slugs
(
  id    SERIAL PRIMARY KEY,
  title TEXT   NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS users
(
  id BIGINT PRIMARY KEY 
);

CREATE TABLE IF NOT EXISTS slugs_users
(
  user_id INTEGER REFERENCES users (id) ON DELETE CASCADE ON UPDATE NO ACTION,
  slug_id INTEGER REFERENCES slugs (id) ON DELETE CASCADE ON UPDATE NO ACTION,
  PRIMARY KEY (user_id, slug_id)
);

CREATE TABLE IF NOT EXISTS records
(
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  slug_title TEXT NOT NULL,
  operation TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  ttl INTERVAL
)