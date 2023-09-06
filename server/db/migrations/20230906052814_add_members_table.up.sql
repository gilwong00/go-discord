CREATE TYPE member_role AS ENUM ('ADMIN','GUEST','MODERATOR');

CREATE TABLE members (
	id uuid PRIMARY KEY,
  role member_role NOT NULL DEFAULT ('GUEST'),
	user_id uuid NOT NULL REFERENCES users (id),
	server_id uuid NOT NULL REFERENCES servers (id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
	updated_at TIMESTAMPTZ NOT NULL
);

CREATE UNIQUE INDEX member_user_id on members USING btree (user_id, server_id);