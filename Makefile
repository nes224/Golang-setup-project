DB_URL=postgresql://root:secret@localhost:5434/bestinter?sslmode=disable

postgres:
	docker run --name postgres13 -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

migratefolder:
	migrate create -ext sql -directory db/migration -seq $(name)

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root bestinter

dropdb:
	docker exec -it postgres13 dropdb bestinter

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/bestinter?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/bestinter?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

.PHONY: postgres migratefolder new_migration createdb dropdb migrateup migratedown sqlc