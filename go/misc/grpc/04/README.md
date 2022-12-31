# Go commands

## Installation

### Server
```
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### Client
```
go get google.golang.org/grpc
```

## Compilation

### Compiling proto stubs

```
protoc ./proto/*.proto --go_out=./server
```

### Compiling GRPC stubs

```
protoc ./proto/*.proto --go_out=./server --go-grpc_out=./server
```

## Links

```
https://developers.google.com/protocol-buffers/docs/overview
https://www.youtube.com/@CodeBangkok/videos
```