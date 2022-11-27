run-pg:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

exec-pg:
	docker exec -it postgres12 psql -U root simple_bank

stop-pg:
	docker stop postgres12
	docker rm postgres12

exec-pg-bash:
	docker exec -it postgres12 /bin/sh

createdb:
	docker exec -t postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -t postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: run-pg exec-pg stop-pg exec-pg-bash createdb dropdb migrateup migratedown test server