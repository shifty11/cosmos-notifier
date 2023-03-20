mock:
    mockgen -source=go/database/telegram_chat_manager.go -destination=go/database/mock_types/telegram_chat_manager.go
    mockgen -source=go/database/discord_channel_manager.go -destination=go/database/mock_types/discord_channel_manager.go

protobuf:
    protoc -I=go/services/grpc/protobuf/ --go_out=go/services/grpc/protobuf/ --go-grpc_out=go/services/grpc/protobuf/ --dart_out=grpc:webapp/lib/api/protobuf/dart/ go/services/grpc/protobuf/*.proto && \
    protoc -I=go/services/grpc/protobuf/ --dart_out=grpc:webapp/lib/api/protobuf/dart/ google/protobuf/timestamp.proto google/protobuf/empty.proto

generate-models:
    cd go && go generate ./ent

generate-migrations:
    cd go && go run main.go database create-migrations

migrate:
    cd go && go run main.go database migrate