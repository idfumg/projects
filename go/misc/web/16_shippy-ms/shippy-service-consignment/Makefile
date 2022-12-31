# export PATH="$PATH:$(go env GOPATH)/bin"
# go install google.golang.org/protobuf/...@v1.25.0
# go get google.golang.org/grpc@v1.32.0
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0.0

compile:
	protoc proto/consignment/*.proto \
		--go_out=paths=source_relative:. -I. \
		--go-grpc_out=paths=source_relative:. -I. \
		--proto_path=.

test:
	go test -race ./...