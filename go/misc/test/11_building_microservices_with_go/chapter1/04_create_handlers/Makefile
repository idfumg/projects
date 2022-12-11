run:
	docker run --rm -v ${PWD}:/app -p 8080:8080 -w /app --name server golang:alpine `go run .`

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o server

build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-s' -o server