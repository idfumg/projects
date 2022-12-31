#!/usr/bin/env bash

export PS4='+(${BASH_SOURCE}:${LINENO}): ${FUNCNAME[0]:+${FUNCNAME[0]}(): }'

set -x

protoc ./proto/*.proto --go_out=./server
protoc ./proto/*.proto --go_out=./server --go-grpc_out=./server
protoc ./proto/*.proto --go_out=./client --go-grpc_out=./client