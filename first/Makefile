.PHONY: migrate-up migrate-down build run

migrate-up:
	migrate -path migrations -database "postgres://go_user:go_pass@localhost:5764/go_db?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://go_user:go_pass@localhost:5764/go_db?sslmode=disable" down

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

build:
	go build -o bin/server .

run: build
	./bin/server