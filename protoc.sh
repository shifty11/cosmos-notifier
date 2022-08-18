#!/bin/sh

protoc -I=go/service_grpc/protobuf/ --go_out=go/service_grpc/protobuf/ --go-grpc_out=go/service_grpc/protobuf/ --dart_out=grpc:webapp/lib/api/protobuf/dart/ go/service_grpc/protobuf/*.proto
protoc -I=go/service_grpc/protobuf/ --dart_out=grpc:webapp/lib/api/protobuf/dart/ google/protobuf/timestamp.proto google/protobuf/empty.proto
