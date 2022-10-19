mock:
	mockgen -source=go/database/telegram_chat_manager.go -destination=go/database/mock_types/telegram_chat_manager.go
	mockgen -source=go/database/discord_channel_manager.go -destination=go/database/mock_types/discord_channel_manager.go

protobuf:
	protoc -I=go/service_grpc/protobuf/ --go_out=go/service_grpc/protobuf/ --go-grpc_out=go/service_grpc/protobuf/ --dart_out=grpc:webapp/lib/api/protobuf/dart/ go/service_grpc/protobuf/*.proto && \
    protoc -I=go/service_grpc/protobuf/ --dart_out=grpc:webapp/lib/api/protobuf/dart/ google/protobuf/timestamp.proto google/protobuf/empty.proto

models:
	cd go && go generate ./ent

migration:
	cd go && go run main.go create-migrations
