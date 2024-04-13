CREATE TABLE IF NOT EXISTS users (
  id bigserial,
  email text NOT NULL UNIQUE DEFAULT '',
  password text NOT NULL DEFAULT '', -- bcrypt hash and salt
  PRIMARY KEY(id)
);
