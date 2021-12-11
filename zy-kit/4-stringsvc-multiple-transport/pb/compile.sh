#!/bin/bash

# Install protc Go:
# go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
# protoc --go_out=. ./svc.proto


# Ref: https://grpc.io/docs/languages/go/basics/
protoc --go_out=. --go_opt=paths=source_relative \                                                                      ─╯
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./svc.proto