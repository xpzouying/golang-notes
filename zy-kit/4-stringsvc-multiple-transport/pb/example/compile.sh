#!/bin/bash

# Install protc Go:
# go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
protoc --go_out=. ./svc.proto