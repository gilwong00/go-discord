.PHONY: start
start:
	cd server && go run cmd/main.go

.PHONY: generate
generate:
	buf generate