DB_DSN := "postgres://dr0zd:secret@localhost:5432/pet_users?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-force:
	$(MIGRATE) force ${VERSION}

version:
	$(MIGRATE) version

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-delete:
	migrate-down 1
	rm -f ./migrations/${NAME}.up.sql ./migrations/${NAME}.down.sql

run:
	go run cmd/server/main.go

lint:
	golangci-lint cache clean
	golangci-lint run --timeout=5m --out-format=colored-line-number