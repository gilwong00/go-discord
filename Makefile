DB_URL=postgres://postgres:postgres@localhost:5432/go_discord?sslmode=disable

.PHONY: start
start:
	cd server && go run cmd/main.go

.PHONY: generate
generate:
	buf generate

.PHONY: startweb
startweb:
	cd client && pnpm run dev

.PHONY: migration
migration:
	cd server && migrate create -ext sql -dir db/migrations

.PHONY: migrateup
migrateup:
	cd server && migrate -path "db/migrations" -database "$(DB_URL)" up

.PHONY: sqlc
sqlc:
	cd server && sqlc generate

.PHONY: dockerup
dockerup:
	cd server/docker/scripts && sh start.sh

.PHONY: dockerdown
dockerdown:
	cd server/docker/scripts && sh destroy.sh