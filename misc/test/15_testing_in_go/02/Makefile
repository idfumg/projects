tests:
	go test -v ./...

tests-all:
	go test -v -tags=integration ./...

web:
	go run ./cmd/web

api:
	go run ./cmd/api

auth:
	curl POST http://localhost:8090/auth -v -j email=admin@example.com password=secret