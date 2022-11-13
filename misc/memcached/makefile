migrate:
	go mod download github.com/golang-migrate/migrate/v4 && \
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate && \
	migrate -path db/migrations/ -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up

services-up:
	docker run \
	-d \
	-e POSTGRES_HOST_AUTH_METHOD=trust \
	-e POSTGRES_USER=user \
	-e POSTGRES_PASSWORD=password \
	-e POSTGRES_DB=dbname \
	-p 5432:5432 \
	--name pg_instance \
	postgres:12.5-alpine && \
	docker run \
	-d \
	-p 11211:11211 \
	--name memcached_instance \
	memcached:1.6.9-alpine

services-down:
	docker stop pg_instance memcached_instance
	docker rm pg_instance memcached_instance

load-test-data:
	go mod download github.com/MarioCarrion/complex-pipelines/part5 && \
	go install github.com/MarioCarrion/complex-pipelines/part5 && \
	DATABASE_URL="postgres://user:password@localhost:5432/dbname?sslmode=disable" part5

run:
	DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable \
	MEMCACHED=localhost:11211 \
	go run .
