Go gRPC Services Course
=======================

PACKAGES
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
go install github.com/golang/mock/mockgen
go get github.com/stretchr/testify/assert
go get github.com/jmoiron/sqlx
go get github.com/golang-migrate/migrate/v4
go get github.com/golang-migrate/migrate/v4/database/postgres
go get github.com/golang-migrate/migrate/v4/source/file
go get github.com/lib/pq

GENERATE MOCKS
go generate ./...

RUN TESTS
go test -v ./...

RUN POSTGRES DOCKER CONTAINER
docker run --name rocket-db -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres

STOP POSTGRES DOCKER CONTAINER
docker stop rocket-db

RUN DOCKER COMPOSE
docker-compose up --build

STOP DOCKER COMPOSE
docker-compose down

CONNECT TO POSTGRES IN DOCKER CONTAINER
docker exec -it grpc-microservice-database bash
psql -U postgres

COOMMANDS IN POSTGRES
\dt -- output a list of relations
\d rockets;

INSTALLING PROTOBUF
brew install protoc-gen-go
brew install protoc-gen-go-grpc
brew install protobuf
