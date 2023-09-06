CREATE TABLE users (
	id uuid PRIMARY KEY,
  username VARCHAR NOT NULL,
	email VARCHAR NOT NULL,
	password VARCHAR NOT NULL,
	avatar_url TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
	deleted_at TIMESTAMPTZ NOT NULL
);
