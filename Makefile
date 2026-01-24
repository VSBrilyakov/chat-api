run:
	docker-compose up -d

test:
	go test -v ./...

migrate:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	goose -dir ./migrations postgres "postgresql://admin:pgpassword123@localhost:5432/chatdb?sslmode=disable" up