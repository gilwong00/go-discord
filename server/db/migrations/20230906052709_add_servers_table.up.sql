CREATE TABLE servers (
	id uuid PRIMARY KEY,
  name VARCHAR NOT NULL,
	image_url TEXT,
	invite_code TEXT UNIQUE,
	created_by_user_id uuid REFERENCES users (id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
	updated_at TIMESTAMPTZ NOT NULL
);

CREATE UNIQUE INDEX server_create_user_id on servers USING btree (created_by_user_id);
