createDb:
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres:17-alpine

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/importer_api?sslmode=disable" -verbose up
migrationdrop:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/importer_api?sslmode=disable" -verbose drop

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createDb postgres migrateup migrationdrop test server