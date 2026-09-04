.PHONY: migrate-up migrate-down build run

migrate-up:
	migrate -path migrations -database "postgres://:@localhost:/?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://:@localhost:/?sslmode=disable" down

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

build:
	go build -o bin/auth .

run: build
	./bin/auth